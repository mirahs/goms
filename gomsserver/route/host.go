package route

import (
	"github.com/gin-gonic/gin"
	"gomsserver/api"
	"gomsserver/middleware"
)


func initHost(group *gin.RouterGroup) {
	host := group.Group("hosts")
	host.Use(middleware.Auth())
	{
		host.GET("", api.HostGetAll)
		host.POST("", api.HostAdd)
		host.PUT(":id", api.HostEdit)
		host.DELETE(":id", api.HostDelete)
	}
}
