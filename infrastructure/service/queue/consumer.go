package queue

import (
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"backend-nabati/infrastructure/shared/constant"
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

func (q queueService) ConsumeData(ctx context.Context, topic string) (err error) {

	cfg := q.rabbitmq.GetConfig()
	notify := cfg.Conn.NotifyClose(make(chan *amqp.Error)) //error channel

	ch, err := cfg.Conn.Channel()
	if err != nil {
		err = Error.New(constant.CONSUMER_BILLING_RABBITMQ, constant.ErrDefineChannelToBroker, err)
		return
	}

	err = ch.ExchangeDeclare(
		topic,                           // name
		constant.RABBITMQ_EXCHANGE_TYPE, // type
		true,                            // durable
		false,                           // auto-deleted
		false,                           // internal
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		return
	}

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		err = Error.New(constant.CONSUMER_BILLING_RABBITMQ, constant.ErrCreateQueueToBroker, err)
		return
	}

	err = cfg.Channel.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		topic,      // exchange
		false,
		nil,
	)
	if err != nil {
		err = Error.New(constant.CONSUMER_BILLING_RABBITMQ, constant.ErrBindingQueueToBroker, err)
		return
	}

	// consume
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		err = Error.New(constant.CONSUMER_BILLING_RABBITMQ, constant.ErrConsumeQueueToBroker, err)
		return
	}

	fmt.Println(fmt.Sprintf(constant.START_LISTENING_TOPIC_FROM_BROKER, topic))
	logger.LogInfo(constant.CONSUMER_BILLING_RABBITMQ, fmt.Sprintf(constant.START_LISTENING_TOPIC_FROM_BROKER, topic))

	for {
		select {
		case err = <-notify:
			if err != nil {
				err = Error.New(constant.CONSUMER_BILLING_RABBITMQ, constant.ErrConsumeQueueToBroker, err)
				for {
					err = q.rabbitmq.Reconnect()
					if err == nil {
						break
					}
				}
			}
		case msg := <-msgs:

			fmt.Println(q.cfg.ProductInsertConsumerName)
			if msg.RoutingKey == q.cfg.ProductInsertConsumerName || msg.Exchange == q.cfg.ProductInsertConsumerName {
				logger.LogInfo(constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))

				fmt.Println(fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))
			} else if msg.RoutingKey == q.cfg.ProductUpdateConsumerName || msg.Exchange == q.cfg.ProductUpdateConsumerName {
				logger.LogInfo(constant.CONSUMER_PRODUCT_UPDATE_RABBITMQ, fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))

				fmt.Println(fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))

			} else if msg.RoutingKey == q.cfg.BillingConsumerName || msg.Exchange == q.cfg.BillingConsumerName {
				logger.LogInfo(constant.CONSUMER_BILLING_RABBITMQ, fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))

				fmt.Println(fmt.Sprintf(constant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))

			}
		}
	}
}
