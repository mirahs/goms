package menu

import (
	"gomsserver/constant"
	"net/http"
	"testing"
)


func Test(t *testing.T) {
	println(Check("/v1/auth/login", http.MethodGet, constant.AdminUserTypeAdmin)) //true
	println(Check("/v1/auth/login", http.MethodGet, constant.AdminUserTypeCS)) //true
	println(Check("/v1/auth/login_token", http.MethodGet, constant.AdminUserTypeAdmin)) //true
	println()

	println(Check("/v1/adm_users", http.MethodGet, constant.AdminUserTypeAdmin)) //true
	println(Check("/v1/adm_users", http.MethodGet, constant.AdminUserTypeCS)) //false
	println()

	println(Check("/v1/adm_users/:id/lock", http.MethodPatch, constant.AdminUserTypeAdmin)) //true
	println(Check("/v1/adm_users/:id/lock", http.MethodPatch, constant.AdminUserTypeCS)) //false
}
