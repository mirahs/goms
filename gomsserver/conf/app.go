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
}


var App *app


// 加载 app 配置文件
func loadApp(mode string)  {
	App = new(app)

	err := ini.MapTo(App, "./conf/app." + mode + ".ini")
	if err != nil {
		panic(err)
	}
}
