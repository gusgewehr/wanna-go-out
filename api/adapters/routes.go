package adapters

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"wanna-go-out-api/infrastructure"
	"wanna-go-out-api/repository"
	"wanna-go-out-api/usecase"
)

func Setup(router *gin.Engine, app *infrastructure.Application) {

	router.Use(cors.Default())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Method not found"})
	})

	apiV1 := router.Group("/api/v1")

	mur := repository.NewMessageRepository(app.Db, app.Logger)

	muc := usecase.NewMessageUseCase(app.Logger, mur)

	mc := &MessageController{muc}

	apiV1.POST("/message", mc.CreateMessageHandler)
	apiV1.GET("/message", mc.GetMessagesHandler)

}
