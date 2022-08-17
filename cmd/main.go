package main

import (
	"backend-nabati/delivery/container"
	"backend-nabati/delivery/http"
	"fmt"
)

func main() {
	container := container.SetupContainer()
	handler := http.SetupHandler(container)

	http := http.ServerHttp(handler)
	http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
