package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func RespSuccess(ctx *gin.Context, data interface{})  {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}

func RespFailed(ctx *gin.Context, httpStatus int, data interface{})  {
	ctx.JSON(httpStatus, data)
}
