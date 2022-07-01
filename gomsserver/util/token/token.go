package token

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)


var tokens sync.Map


/**
@description 生成 token
@param id uint32 用户 id
@param multi bool 是否可以多终端登录
 */
func Make(id uint32, multi bool) string {
	if !multi {
		destroyById(id)
	}
	token := fmt.Sprintf("%s", uuid.NewV4())
	tokens.Store(token, id)
	return token
}

/**
@description 获取用户 id
@param token string
 */
func Get(token string) uint32 {
	val, ok := tokens.Load(token)
	if ok {
		return val.(uint32)
	}
	return 0
}

/**
@description 注销token
@param token string 要注销的 token
 */
func Destroy(token string) {
	tokens.Delete(token)
}


func destroyById(id uint32) {
	var delTokens []string
	tokens.Range(func (key, val interface{}) bool {
		if id == val {
			delTokens = append(delTokens, key.(string))
		}
		return true
	})

	for _, token := range delTokens {
		Destroy(token)
	}
}


func printAll() {
	tokens.Range(func (key, val interface{}) bool {
		fmt.Printf("printAll key:%s, val:%d\n", key, val)
		return true
	})
}
