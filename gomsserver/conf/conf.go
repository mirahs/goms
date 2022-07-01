package conf


func init()  {
	nodesInit()
}


// Load 加载配置文件
func Load(mode string)  {
	loadApp(mode)
}
