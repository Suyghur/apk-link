/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitSdkRouter(r *gin.RouterGroup) {
	sdkRouter := r.Group("sdk")
	{
		fuseRouter := sdkRouter.Group("fuse").Use(middleware.JWTAuth())
		{
			fuseRouter.GET("getFuseSdks", v1.GetFuseSdks)
			fuseRouter.GET("searchFuseSdk", v1.SearchFuseSdk)
			fuseRouter.PUT("modifyFuseSdk", v1.ModifyFuseSdk)
			fuseRouter.POST("createFuseSdk", v1.CreateFuseSdk)

		}
		channelRouter := sdkRouter.Group("channel").Use(middleware.JWTAuth())
		{
			channelRouter.GET("getChannelSdks", v1.GetChannelSdks)
			channelRouter.GET("searchChannelSdk", v1.SearchChannelSdk)
			channelRouter.PUT("modifyChannelSdk", v1.ModifyChannelSdk)
			channelRouter.POST("createChannelSdk", v1.CreateChannelSdk)
		}
		pluginRouter := sdkRouter.Group("plugin").Use(middleware.JWTAuth())
		{
			pluginRouter.GET("getPluginSdks", v1.GetPluginSdks)
			pluginRouter.GET("searchPluginSdk", v1.SearchPluginSdk)
			pluginRouter.PUT("modifyPluginSdk", v1.ModifyPluginSdk)
			pluginRouter.POST("createPluginSdk", v1.CreatePluginSdk)
		}
	}
}
