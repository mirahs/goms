package api

import (
	"github.com/gin-gonic/gin"
	"gomsserver/common"
	"gomsserver/dao"
)


// CSServerHostHeart host 心跳
func CSServerHostHeart(ctx *gin.Context) {
	host := ctx.Params.ByName("host")

	dao.HostOnlineHeart(host)

	common.RespSuccess(ctx, nil)
}
