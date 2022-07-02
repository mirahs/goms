package service

import (
	"github.com/tidwall/gjson"
	"gomsserver/conf"
	"gomsserver/util"
	"log"
)


func HostClientHeart() {
	url := conf.App.ServerApi + "/cs_servers/host_heart/" + conf.NodeName

	resp, err := util.ReqGet(url)
	if err != nil {
		log.Printf("service.HostClientHeart ReqGet err:%v\n", err)
		return
	}

	respStr := resp.String()

	code := gjson.Get(respStr, "code").Int()
	if code != 0 {
		msg := gjson.Get(respStr, "msg").String()
		log.Printf("service.HostClientHeart resp msg:%s\n", msg)
		return
	}
}
