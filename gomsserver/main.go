package main

import (
	"flag"
	"fmt"
	"gomsserver/conf"
	"gomsserver/model"
	"gomsserver/route"
	"gomsserver/util"
)


func main() {
	mode := flag.String("mode", "dev", "运行环境 dev|test|prod")
	flag.Parse()

	conf.Load(*mode)
	model.Init()
	defer model.Close()
	util.IpInit(conf.App.Ip2RegionDbFile)
	defer util.IpClose()

	engine := route.Init()

	panic(engine.Run(fmt.Sprintf(":%d", conf.App.GinPort)))
}
