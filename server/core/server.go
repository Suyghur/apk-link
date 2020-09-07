/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : server
*/
package core

import (
	"fmt"
	"server/global"
	"server/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		//初始化redis服务
		initialize.Redis()
	}
	r := initialize.Routers()
	r.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	address = "0.0.0.0" + address
	//s := initServer(address, r)
	//保障文本顺序输出
	time.Sleep(10 * time.Millisecond)
	global.GVA_LOG.Debug("server run success on ", address)

	global.GVA_LOG.Info("欢迎使用 apk-link")

	//global.GVA_LOG.Error(s.ListenAndServe())

	//windows环境用这个方法启动
	r.Run(address)
}

//非Windows环境解开
//func initServer(address string, r *gin.Engine) server {
//	s := endless.NewServer(address, r)
//	s.ReadHeaderTimeout = 10 * time.Millisecond
//	s.WriteTimeout = 10 * time.Second
//	s.MaxHeaderBytes = 1 << 20
//	return s
//}
