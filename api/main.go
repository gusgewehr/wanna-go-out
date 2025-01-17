package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"wanna-go-out-api/adapters"
	"wanna-go-out-api/infrastructure"
)

func main() {

	envPath := ".env"

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	app := infrastructure.App(envPath)

	httpServer := gin.Default()

	adapters.Setup(httpServer, app)

	err := httpServer.Run()

	if err != nil {
		panic(err)
	}

}
