package route

import (
	"github.com/gin-gonic/gin"
	"gomsserver/api"
)


func initCSServer(group *gin.RouterGroup) {
	server := group.Group("cs_servers")

	server.GET("host_heart/:host", api.CSServerHostHeart)
}
