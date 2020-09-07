package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	initialize.Mysql()
	initialize.DBTables()
	//结束前关闭数据块链接
	defer global.GVA_DB.Close()
	core.RunServer()
}
