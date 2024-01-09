package cmd

import (
	"backend-nabati/delivery/container"
	"backend-nabati/delivery/http"
	"backend-nabati/delivery/queue"
	"backend-nabati/infrastructure/logger"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
)

func Execute() {
	// start init container
	container := container.SetupContainer()

	// start queue service
	queue.StartQueueServices(container)

	// start http service
	http := http.ServeHttp(container)
	err := http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
	if err != nil {
		// Handle the error, e.g., log it or exit the application
		fmt.Printf("Error starting HTTP server: %s\n", err)
		// Optionally, you can log the error or exit the application.

		logger.LogError(constant.ErrHttpServer, constant.ErrHttpServer, err.Error())

	}
}
