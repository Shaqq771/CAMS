package rabbitmq

import (
	"backend-nabati/infrastructure/shared/constant"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitmqConfig struct {
	Host                      string
	Username                  string
	Password                  string
	Port                      int
	BillingProducerName       string
	BillingConsumerName       string
	ProductInsertConsumerName string
	ProductUpdateConsumerName string
}

type RabbitMQ interface {
	Connect() (err error)
	Close()
	Reconnect() error
	GetConfig() rabbitMQ
}

type rabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Err     chan error
	config  RabbitmqConfig
}

func NewConnection(config RabbitmqConfig) RabbitMQ {
	return &rabbitMQ{
		config: config,
		Err:    make(chan error),
	}
}

func (c *rabbitMQ) GetConfig() rabbitMQ {
	return *c
}

func (c *rabbitMQ) Connect() (err error) {
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

	if c.config.Port == 0 {
		connPattern = "amqp://%v:%v@%v"
		clientUrl = fmt.Sprintf(connPattern,
			c.config.Username,
			c.config.Password,
			c.config.Host,
		)
	}

	c.Conn, err = amqp.Dial(clientUrl)
	if err != nil {
		fmt.Println(err)
		if err = c.Retry(); err != nil {
			err = fmt.Errorf(constant.ErrConnectToBroker, err)
			return
		}
	}

	c.Channel, err = c.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf(constant.ErrCreateChannelToBroker, err)
		return
	}

	if err = c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		err = fmt.Errorf(constant.ErrSetupQueueToBroker, err)
		return
	}

	return
}

func (c *rabbitMQ) Retry() (err error) {
	fmt.Println(constant.RETRY_MESSAGE_BROKER)
	connPattern := "amqp://%v:%v@%v:%v"
	if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		c.config.Username,
		c.config.Password,
		c.config.Host,
		constant.RABBITMQ_PORT,
	)

	if c.config.Port == 0 {
		connPattern = "amqp://%v:%v@%v"
		clientUrl = fmt.Sprintf(connPattern,
			c.config.Username,
			c.config.Password,
			c.config.Host,
		)
	}

	conn, err := amqp.Dial(clientUrl)
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToBroker, err)
		return
	}

	c.Conn = conn

	return
}

func (c *rabbitMQ) Close() {
	c.Conn.Close()
}

func (c *rabbitMQ) Reconnect() error {
	fmt.Println(constant.RETRY_MESSAGE_BROKER)
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}
