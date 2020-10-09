package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	initialize.Gorm()
	if global.GvaConfig.System.NeedInitData {

	}
	//结束前关闭数据块链接
	db, _ := global.GvaDb.DB()
	defer db.Close()
	core.RunServer()
}
