package rabbitmq

import (
	"backend-nabati/infrastructure/shared/constant"
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

func (c *rabbitMQ) Consume(context context.Context, topic string) (msgs <-chan amqp.Delivery, err error) {

	select {
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		err = fmt.Errorf(constant.ErrCreateChannelToBroker, err)
		return
	}

	// declare queue
	// if queue not exist will be created
	queue, err := c.channel.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	// Handle any errors if we were unable to create the queue
	if err != nil {
		err = fmt.Errorf(constant.ErrCreateQueueToBroker, err)
		return
	}

	err = c.channel.QueueBind(
		queue.Name, // queue name
		topic,      // routing key
		"tasks",    // exchange
		false,
		nil,
	)
	if err != nil {
		err = fmt.Errorf(constant.ErrBindingQueueToBroker, err)
		return
	}

	// consume
	msgs, err = c.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		err = fmt.Errorf(constant.ErrConsumeQueueToBroker, err)
		return
	}

	return
}
