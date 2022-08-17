package rabbitmq

import (
	"backend-nabati/infrastructure/shared/constant"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func (c *Connection) Publish(context context.Context, routingKey string, event interface{}) (status bool, err error) {

	select {
	case err := <-c.err:
		if err != nil {
			fmt.Println(c.err)
			c.Reconnect()
		}
	default:
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		err = fmt.Errorf(constant.ErrCreateChannelToBroker, err)
		return
	}

	body, err := json.Marshal(event)
	if err != nil {
		err = fmt.Errorf(constant.ErrToMarshalJSON, err)
	}

	err = c.channel.Publish(
		"tasks",    // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json", // XXX: We will revisit this in future episodes
			Body:        body,
			Timestamp:   time.Now(),
		})
	if err != nil {
		err = fmt.Errorf(constant.ErrPublishQueueToBroker, err)
		return
	}

	fmt.Println("Successfully Published Message to Queue")

	return true, nil
}
