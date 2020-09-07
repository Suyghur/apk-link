/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_channel
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
	"linking-api/middleware"
)

func InitChannelRouter(r *gin.RouterGroup) {
	channelRouter := r.Group("channel").Use(middleware.JWTAuth())
	{
		channelRouter.POST("createChannel", v1.CreateChannel)
		channelRouter.POST("delChannel", v1.DelChannel)
		channelRouter.POST("modifyChannel", v1.ModifyChannel)
		channelRouter.POST("listChannel", v1.ListChannel)
	}
}
