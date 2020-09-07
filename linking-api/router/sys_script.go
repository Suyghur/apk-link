/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_script
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
	"linking-api/middleware"
)

func InitScriptRouter(r *gin.RouterGroup) {
	scriptRouter := r.Group("script").Use(middleware.JWTAuth())
	{
		scriptRouter.POST("list", v1.ListScript)
		//scriptRouter.POST("search", v1.SearchScript)
		scriptRouter.POST("modify", v1.ModifyScript)
		scriptRouter.POST("create", v1.CreateScript)

		//scriptRouter.POST("create", v1.ModifyScript)
		//gameScriptRouter := scriptRouter.Group("game")
		//{
		//	gameScriptRouter.POST("list", v1.ListScript)
		//	gameScriptRouter.POST("search", v1.SearchScript)
		//	gameScriptRouter.POST("modify", v1.ModifyGameScript)
		//	gameScriptRouter.POST("")
		//}
	}

}
