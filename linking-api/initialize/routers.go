/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : routers
*/
package initialize

import (
	"github.com/gin-gonic/gin"
	"linking-api/global"
	"linking-api/middleware"
	"linking-api/router"
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
	router.InitUserRouter(apiGroup) //注册用户路由
	//router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	//router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	//router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	//router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	//router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	//router.InitSimpleUploaderRouter(ApiGroup)        // 断点续传（插件版）
	//router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	//router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	//router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	//router.InitSystemRouter(ApiGroup)                // system相关路由
	//router.InitCustomerRouter(ApiGroup)              // 客户路由
	//router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	//router.InitSysDictionaryDetailRouter(ApiGroup)   // 字典详情管理
	//router.InitSysDictionaryRouter(ApiGroup)         // 字典管理
	//router.InitSysOperationRecordRouter(ApiGroup)    // 操作记录
	global.GVA_LOG.Info("router register success")
	return r
}
