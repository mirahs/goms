package common

import "github.com/gin-gonic/gin"


var (
	// 0 - 199 系统(0 表示成功)
	ErrTokenInvalid			= gin.H{"code": 1, "msg": "登录过期, 请重新登录"}
	ErrParamFailed			= gin.H{"code": 2, "msg": "参数错误"}
	ErrTokenDeniedEmpty 	= gin.H{"code": 3, "msg": "权限不足"}
	ErrTokenDeniedNoUser	= gin.H{"code": 4, "msg": "权限不足"}
	ErrDb                	= gin.H{"code": 5, "msg": "内部错误"}
	ErrDbTransaction     	= gin.H{"code": 12, "msg": "内部错误"}

	// 200 - 299 adm_user
	ErrAUAccountNotExist 	= gin.H{"code": 200, "msg": "账号不存在"}
	ErrAUPasswordNotSame 	= gin.H{"code": 201, "msg": "密码错误"}
	ErrAUOperateSelf     	= gin.H{"code": 202, "msg": "不能操作自己"}
	ErrAUIdNotExist      	= gin.H{"code": 203, "msg": "账号不存在"}
	ErrAUAccountExist 		= gin.H{"code": 204, "msg": "账号已存在"}
	ErrAULocked 			= gin.H{"code": 205, "msg": "账号已封禁"}
	ErrAUDenied 			= gin.H{"code": 206, "msg": "权限不足"}
	ErrAUPasswordFailedMax 	= gin.H{"code": 207, "msg": "密码错误次数过多"}

	// 300 - 319 host
	ErrHostExist 			= gin.H{"code": 300, "msg": "主机已存在"}
	ErrHostIdNotExist      	= gin.H{"code": 301, "msg": "主机不存在"}
)
