package middleware

import (
	"github.com/gin-gonic/gin"
	"gomsserver/common"
	"gomsserver/constant"
	"gomsserver/menu"
	"gomsserver/model"
	"gomsserver/util/token"
	"net/http"
)


func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(constant.KeyHttpHeaderXToken)
		if tokenString == "" {
			common.RespFailed(ctx, http.StatusBadRequest, common.ErrTokenDeniedEmpty)
			ctx.Abort()
			return
		}

		userId := token.Get(tokenString)
		if userId == 0 {
			common.RespFailed(ctx, http.StatusNotFound, common.ErrTokenInvalid)
			ctx.Abort()
			return
		}

		user, err := model.AUGet(userId)
		if err != nil {
			common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
			ctx.Abort()
			return
		}
		if user.ID == 0 {
			common.RespFailed(ctx, http.StatusNotFound, common.ErrTokenDeniedNoUser)
			ctx.Abort()
			return
		}
		if user.Locked == 1 {
			common.RespFailed(ctx, http.StatusNotFound, common.ErrAULocked)
			ctx.Abort()
			return
		}

		if !menu.Check(ctx.FullPath(), ctx.Request.Method, user.Type) {
			common.RespFailed(ctx, http.StatusNotFound, common.ErrAUDenied)
			ctx.Abort()
			return
		}

		ctx.Set(constant.KeyAdmUserId, userId)
		ctx.Set(constant.KeyAdmUser, user)
		ctx.Next()
	}
}
