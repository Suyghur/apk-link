/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : routers
*/
package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/middleware"
	"server/router"
)

func Routers() *gin.Engine {
	var r = gin.Default()
	//开启https
	//router.User(middleware.LoadTls())
	global.GVA_LOG.Debug("use middleware logger")

	//跨域
	r.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware cors")

	//swagger
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Debug("register swagger handler")

	//方便统一添加路由组前缀 多服务器上线使用
	apiGroup := r.Group("v1")
	router.InitBaseRouter(apiGroup)
	router.InitUserRouter(apiGroup) //注册用户路由
	router.InitTaskRouter(apiGroup)
	router.InitSdkRouter(apiGroup)
	router.InitKeystoreRouter(apiGroup)
	router.InitScriptRouter(apiGroup)
	router.InitOriginRouter(apiGroup)
	router.InitOptionsRouter(apiGroup)
	router.InitGameRouter(apiGroup)
	router.InitChannelRouter(apiGroup)
	router.InitPluginRouter(apiGroup)
	router.InitLinkRouter(apiGroup)
	global.GVA_LOG.Info("router register success")
	return r
}
