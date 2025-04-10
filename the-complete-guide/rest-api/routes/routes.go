package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/htutwaiphyoe/mastering-go/the-complete-guide/api/controllers"
	"github.com/htutwaiphyoe/mastering-go/the-complete-guide/api/middlewares"
)

func Register(server *gin.Engine) {

	authRoutes := server.Group("/")
	authRoutes.Use(middlewares.Authenticate)
	authRoutes.GET("/events", controllers.GetEvents)
}
