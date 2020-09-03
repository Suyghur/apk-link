/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
)

func InitSdkRouter(r *gin.RouterGroup) {
	sdkRouter := r.Group("sdk")
	{
		fuseRouter := sdkRouter.Group("fuse")
		{
			fuseRouter.POST("list", v1.ListFuseSdk)
			fuseRouter.POST("modify", v1.ModifyFuseSdk)
			fuseRouter.POST("create", v1.CreateFuseSdk)

		}
		channelRouter := sdkRouter.Group("channel")
		{
			channelRouter.POST("list", v1.ListChannelSdk)
			channelRouter.POST("modify", v1.ModifyChannelSdk)
			channelRouter.POST("create", v1.CreateChannelSdk)
		}
		pluginRouter := sdkRouter.Group("plugin")
		{
			pluginRouter.POST("list", v1.ListPluginSdk)
			pluginRouter.POST("modify", v1.ModifyPluginSdk)
			pluginRouter.POST("create", v1.CreatePluginSdk)
		}
		//sdkRouter.POST("fuse/list", v1.ListFuseSdk)
		//sdkRouter.POST("channel/list")
		//sdkRouter.POST("plugin/list")
	}
}
