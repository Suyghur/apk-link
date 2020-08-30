/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_script
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
)

func InitScriptRouter(r *gin.RouterGroup) {
	scriptRouter := r.Group("script")
	{
		gameScriptRouter := scriptRouter.Group("game")
		{
			gameScriptRouter.POST("list", v1.ListGameScript)
			gameScriptRouter.POST("search", v1.SearchGameScript)
			gameScriptRouter.POST("modify", v1.ModifyGameScript)
		}
	}

}
