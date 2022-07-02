package route

import (
	"github.com/gin-gonic/gin"
	"gomsserver/conf"
	"gomsserver/middleware"
)


func Init() *gin.Engine {
	gin.SetMode(conf.App.GinMode)

	var engine *gin.Engine
	if conf.App.GinMode == gin.DebugMode {
		engine = gin.Default()
	} else {
		engine = gin.New()
		engine.Use(gin.Recovery())
	}

	initRoute(engine)

	return engine
}


func initRoute(engine *gin.Engine) {
	engine.Use(middleware.Cors())

	v1 := engine.Group("v1")

	if conf.NodeType == "server" {
		initAuth(v1)

		initAdmUser(v1)
		initLog(v1)

		initHost(v1)

		initCSServer(v1)
	}
}
