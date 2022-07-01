package main

import (
	"flag"
	"fmt"
	"gomsserver/conf"
	"gomsserver/model"
	"gomsserver/route"
	"gomsserver/util"
	"os"
)


func main() {
	mode := flag.String("mode", "dev", "运行环境 dev|test|prod")
	nodeType := flag.String("t", "", "节点类型 server|client")
	nodeName := flag.String("n", "", "节点名称")
	flag.Parse()

	if *nodeType == "" || *nodeName == "" {
		flag.Usage()
		os.Exit(-1)
	}

	conf.NodeType = *nodeType
	conf.NodeName = *nodeName

	conf.Load(*mode)

	if *nodeType == "server" {
		model.Init()
		defer model.Close()
		util.IpInit(conf.App.Ip2RegionDbFile)
		defer util.IpClose()
	}

	engine := route.Init()

	panic(engine.Run(fmt.Sprintf(":%d", conf.App.GinPort)))
}
