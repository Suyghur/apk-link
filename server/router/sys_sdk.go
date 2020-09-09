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
			fuseRouter.POST("getFuseSdks", v1.GetFuseSdks)
			fuseRouter.POST("searchFuseSdk", v1.SearchFuseSdk)
			fuseRouter.POST("modifyFuseSdk", v1.ModifyFuseSdk)
			fuseRouter.POST("createFuseSdk", v1.CreateFuseSdk)

		}
		channelRouter := sdkRouter.Group("channel").Use(middleware.JWTAuth())
		{
			channelRouter.POST("getChannelSdks", v1.GetChannelSdks)
			channelRouter.POST("searchChannelSdk", v1.SearchChannelSdk)
			channelRouter.POST("modifyChannelSdk", v1.ModifyChannelSdk)
			channelRouter.POST("createChannelSdk", v1.CreateChannelSdk)
		}
		pluginRouter := sdkRouter.Group("plugin").Use(middleware.JWTAuth())
		{
			pluginRouter.POST("getPluginSdks", v1.GetPluginSdks)
			pluginRouter.POST("searchPluginSdk", v1.SearchPluginSdk)
			pluginRouter.POST("modifyPluginSdk", v1.ModifyPluginSdk)
			pluginRouter.POST("createPluginSdk", v1.CreatePluginSdk)
		}
		//sdkRouter.POST("fuse/list", v1.SearchFuseSdk)
		//sdkRouter.POST("channel/list")
		//sdkRouter.POST("plugin/list")
	}
}
