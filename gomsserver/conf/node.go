package conf


var NodeType string
var NodeName string


// 节点通讯密钥
var nodes = make(map[string]string)


func nodesInit() {
	nodes["127.0.0.1"] = "dfeafaff"
}
