package consumer

import (
	"backend-nabati/domain/sales/feature"
	"backend-nabati/domain/shared/context"
	"backend-nabati/infrastructure/broker/rabbitmq"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
)

type SalesConsumer interface {
	Consumer() error
}

type salesConsumer struct {
	rabbitmq     rabbitmq.RabbitMQ
	salesFeature feature.SalesFeature
}

func NewSalesConsumer(rabbitmq rabbitmq.RabbitMQ, salesFeature feature.SalesFeature) SalesConsumer {
	return &salesConsumer{
		rabbitmq:     rabbitmq,
		salesFeature: salesFeature,
	}
}

func (sc salesConsumer) Consumer() error {
	ctx := context.CreateContext()

	go func() {
		msgs, err := sc.rabbitmq.Consume(ctx, constant.SalesTopic)
		if err != nil {
			return
		}

		forever := make(chan bool)
		go func() {
			for msg := range msgs {
				fmt.Println("received message from:", msg.Exchange)

				switch msg.RoutingKey {
				case constant.SalesTopic:
					// Update
					err = sc.salesFeature.UpdateSalesProductFromBroker(string(msg.Body))
					if err != nil {
					}
				default:
				}
			}
		}()

		fmt.Println(" [*] - Waiting for messages from", constant.SalesTopic)
		<-forever

	}()

	return nil
}
