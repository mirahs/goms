package conf

import "gopkg.in/ini.v1"


type app struct {
	GinMode string //gin 模式(gin.DebugMode|gin.ReleaseMode|gin.TestMode)
	GinPort int

	MysqlHost     string
	MysqlPort     int
	MysqlDatabase string
	MysqlUser     string
	MysqlPassword string
	MysqlCharset  string

	InitAccount     string //后台初始账号
	InitPassword    string //后台初始密码

	Ip2RegionDbFile string //Ip2Region 数据文件

	ServerApi	string //server api 地址(client 需要配置)
	ClientPort	string //client 端口(server 需要配置)
}


var App *app


// 加载 app 配置文件
func loadApp(mode string)  {
	App = new(app)

	err := ini.MapTo(App, "./conf/app." + mode + ".ini")
	if err != nil {
		panic(err)
	}

	if NodeType == "client" && NodeName == "127.0.0.1" {
		App.GinPort++
	}
}
