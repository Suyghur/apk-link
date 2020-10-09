/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : server
*/
package core

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GvaConfig.System.UseMultipoint {
		//初始化redis服务
		initialize.Redis()
	}
	r := initialize.Routers()
	r.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GvaConfig.System.Addr)

	address = "0.0.0.0" + address
	s := initServer(address, r)
	//保障文本顺序输出
	time.Sleep(10 * time.Millisecond)
	global.GvaLog.Debug("server run success on ", zap.String("address", address))

	global.GvaLog.Info("欢迎使用 apk-link")

	global.GvaLog.Error(s.ListenAndServe().Error())

	//windows环境用这个方法启动
	//r.Run(address)
}
