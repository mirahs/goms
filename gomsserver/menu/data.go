package menu

import (
	"gomsserver/constant"
	"net/http"
)


type item struct {
	Path string //路径匹配
	Methods []method //http 方法列表
	Types []uint8 //用户类型
}

type method struct {
	Method string //http 方法
	Types []uint8 //用户类型
}


var items []*item


func init() {
	items = []*item{
		{Path: "/v1/auth/login"},
		{Path: "/v1/auth/logout"},
		{Path: "/v1/auth/login_token"},

		{Path: "/v1/adm_users/info"},
		{Path: "/v1/adm_users", Methods: []method{
			{http.MethodGet, []uint8{constant.AdminUserTypeAdmin}},
			{http.MethodPost, []uint8{constant.AdminUserTypeAdmin}},
		},
		},
		{Path: "/v1/adm_users/:id", Methods: []method{
			{http.MethodPut, []uint8{constant.AdminUserTypeAdmin}},
			{http.MethodDelete, []uint8{constant.AdminUserTypeAdmin}},
		},
		},
		{Path: "/v1/adm_users/:id/lock", Types:[]uint8{constant.AdminUserTypeAdmin}},
		{Path: "/v1/adm_users/:id/reset", Types:[]uint8{constant.AdminUserTypeAdmin}},
		{Path: "/v1/adm_users/password"},

		{Path: "/v1/log/login/adm_user", Types: []uint8{constant.AdminUserTypeAdmin}},
		{Path: "/v1/log/login/adm_user/:id", Types: []uint8{constant.AdminUserTypeAdmin}},

		{Path: "/v1/hosts", Types: []uint8{constant.AdminUserTypeAdmin}},
		{Path: "/v1/hosts/:id", Types: []uint8{constant.AdminUserTypeAdmin}},
	}
}
