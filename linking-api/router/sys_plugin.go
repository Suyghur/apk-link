/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_plugin
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
	"linking-api/middleware"
)

func InitPluginRouter(r *gin.RouterGroup) {
	pluginRouter := r.Group("plugin").Use(middleware.JWTAuth())
	{
		pluginRouter.POST("createPlugin", v1.CreatePlugin)
		pluginRouter.POST("delPlugin", v1.DelPlugin)
		pluginRouter.POST("modifyPlugin", v1.ModifyPlugin)
		pluginRouter.POST("listPlugin", v1.ListPlugin)
	}
}
