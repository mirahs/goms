package token

import (
	"testing"
)


// 测试 token 不能多终端登录
func Test(t *testing.T) {
	Make(1, false)
	Make(2, false)
	Make(3, false)
	Make(1, false)

	printAll()
}

// 测试 token 多终端登录
func Test2(t *testing.T) {
	Make(1, true)
	Make(2, true)
	Make(3, true)
	Make(1, true)

	printAll()
}
