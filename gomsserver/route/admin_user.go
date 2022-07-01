package route

import (
	"github.com/gin-gonic/gin"
	"gomsserver/api"
	"gomsserver/middleware"
)


func initAdmUser(group *gin.RouterGroup) {
	user := group.Group("adm_users")
	user.Use(middleware.Auth())
	{
		user.GET("info", api.AdmUserInfo)

		user.GET("", api.AdmUserGetAll)
		user.POST("", api.AdmUserAdd)
		user.PUT(":id", api.AdmUserEdit)
		user.DELETE(":id", api.AdmUserDelete)

		user.PATCH(":id/lock", api.AdmUserLock)
		user.PATCH(":id/reset", api.AdmUserReset)
		user.PATCH("password", api.AdmUserPassword)
	}
}
