package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


var validate = validator.New()


// 绑定并验证
func GinBindValid(ctx *gin.Context, bind interface{}) error {
	err := ctx.ShouldBind(bind)
	if err != nil {
		return err
	}

	err = validate.Struct(bind)

	if err != nil {
		errFirst := err.(validator.ValidationErrors)[0]
		return errors.New(fmt.Sprintf("%s Value:%v", errFirst, errFirst.Value()))
	}

	return nil
}


func GinIP(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("X-Forward-For")
	if ip != "" {
		return ip
	}

	return ctx.ClientIP()
}
