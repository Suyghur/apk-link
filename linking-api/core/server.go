/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : linking-api
*/
package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"linking-api/global"
	"linking-api/initialize"
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
	s := initServer(address, r)
	//保障文本顺序输出
	time.Sleep(10 * time.Millisecond)
	global.GVA_LOG.Debug("linking-api run success on ", address)

	fmt.Printf(`欢迎使用 linking-api`)
	global.GVA_LOG.Error(s.ListenAndServe())
}

func initServer(address string, r *gin.Engine) server {
	s := endless.NewServer(address, r)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
