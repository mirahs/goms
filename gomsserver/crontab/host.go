package crontab

import (
	"gomsserver/conf"
	"gomsserver/service"
)


func hostInitCs() {
	if conf.NodeType == "client" {
		hostInitCsClient()
	}
}


func hostInitCsClient() {
	// 每 5 秒 host 心跳
	_, _ = cs.AddFunc("*/5 * * * * *", func() {
		defer recoverCheck()

		//println()
		//log.Println("每 5 秒 host 心跳")
		service.HostClientHeart()
	})
}
