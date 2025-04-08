package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/htutwaiphyoe/mastering-go/the-complete-guide/api/controllers"
)

func Register(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
}
