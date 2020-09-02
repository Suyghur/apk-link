package main

import (
	"linking-api/core"
	"linking-api/global"
	"linking-api/initialize"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()
	//结束前关闭数据块链接
	defer global.GVA_DB.Close()
	core.RunServer()
}
