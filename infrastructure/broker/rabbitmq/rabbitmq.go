package rabbitmq

import (
	"backend-nabati/infrastructure/shared/constant"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitmqConfig struct {
	Host         string
	Username     string
	Password     string
	Port         int
	ConsumerName string
}

type Connection struct {
	name    string
	conn    *amqp.Connection
	channel *amqp.Channel
	err     chan error
	config  RabbitmqConfig
}

var (
	connectionPool = make(map[string]*Connection)
)

func NewConnection(name string, config RabbitmqConfig) *Connection {
	if c, ok := connectionPool[name]; ok {
		return c
	}
	c := &Connection{
		config: config,
		err:    make(chan error),
	}
	connectionPool[name] = c
	return c
}

func GetConnection(name string) *Connection {
	return connectionPool[name]
}

func (c *Connection) Connect() (err error) {
	connPattern := "amqp://%v:%v@%v:%v"
	if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		c.config.Username,
		c.config.Password,
		c.config.Host,
		c.config.Port,
	)

	c.conn, err = amqp.Dial(clientUrl)
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToBroker, err)
		return
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		err = fmt.Errorf(constant.ErrCreateChannelToBroker, err)
		return
	}

	err = c.channel.ExchangeDeclare(
		"tasks", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		err = fmt.Errorf(constant.ErrCreateTopicToBroker, err)
		return
	}

	if err = c.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		err = fmt.Errorf(constant.ErrSetupQueueToBroker, err)
		return
	}

	return
}

func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}
