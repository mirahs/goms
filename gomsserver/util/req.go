package util

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
)


func ReqGet(url string) (resp *req.Response, err error) {
	resp, err = req.Get(url)
	if err != nil {
		return
	}

	if !resp.IsSuccess() {
		err = errors.New(fmt.Sprintf("IsSuccess false response:%v\n", resp))
		return
	}

	return
}
