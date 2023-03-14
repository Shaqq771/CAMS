package queue

import (
	"backend-nabati/delivery/container"
	"context"
	"fmt"
	"log"
)

func StartQueueServices(cont container.Container) {
	fmt.Println("Starting queue service...")

	ctx := context.Background()

	// Listening billing update
	go func() {
		err := cont.QueueServices.ConsumeData(ctx, cont.EnvironmentConfig.RabbitMq.BillingConsumerName)
		if err != nil {
			log.Panic(err)
		}
	}()

	// // Listening product insert
	go func() {
		err := cont.QueueServices.ConsumeData(ctx, cont.EnvironmentConfig.RabbitMq.ProductInsertConsumerName)
		if err != nil {
			log.Panic(err)
		}

	}()

	// // Listening product update
	go func() {
		err := cont.QueueServices.ConsumeData(ctx, cont.EnvironmentConfig.RabbitMq.ProductUpdateConsumerName)
		if err != nil {
			log.Panic(err)
		}
	}()

}
