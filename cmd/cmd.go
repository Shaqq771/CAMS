package cmd

import (
	"backend-nabati/delivery/container"
	"backend-nabati/delivery/http"
	"backend-nabati/delivery/queue"
	"fmt"
)

func Execute() {
	// start init container
	container := container.SetupContainer()

	// start queue service
	queue.StartQueueServices(container)

	// start http service
	http := http.ServeHttp(container)
	http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
