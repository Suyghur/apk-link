/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_options
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitOptionsRouter(r *gin.RouterGroup) {
	optionsRouter := r.Group("options")
	{
		optionsRouter.GET("getOptions", v1.GetOptions)
		//optionsRouter.GET("getAids",v1.GetAids)
		//optionsRouter.POST("channelName")
		//optionsRouter.POST("pluginName")
	}
}
