package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/htutwaiphyoe/mastering-go/the-complete-guide/api/routes"
)

func Setup() {
	server := gin.Default()

	routes.Register(server)

	server.Run(":8080")
}
