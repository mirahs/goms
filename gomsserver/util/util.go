package util

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)


func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func BCryptHash(plainPwd string) string {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	return string(hashedPwd)
}

func BCryptCompare(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

func InSlice[K comparable](items []K, item K) bool {
	for _, val := range items {
		if val == item {
			return true
		}
	}
	return false
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func RandStr(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}


func PageInfo(ctx *gin.Context) (int, int) {
	page := com.StrTo(ctx.Query("_page")).MustInt()
	pageSize := com.StrTo(ctx.Query("_page_size")).MustInt()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return page, pageSize
}
