/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_script
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitScriptRouter(r *gin.RouterGroup) {
	scriptRouter := r.Group("script").Use(middleware.JWTAuth())
	{
		scriptRouter.GET("searchScript", v1.SearchScript)
		//scriptRouter.POST("search", v1.SearchScript)
		scriptRouter.PUT("modifyScript", v1.ModifyScript)
		scriptRouter.POST("createScript", v1.CreateScript)
	}

}
