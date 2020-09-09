/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_game_group
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitGameRouter(r *gin.RouterGroup) {
	gameRouter := r.Group("game").Use(middleware.JWTAuth())
	{
		gameRouter.POST("createGame", v1.CreateGame)
		gameRouter.POST("deleteGame", v1.DelGame)
		gameRouter.POST("modifyGame", v1.ModifyGame)
		gameRouter.POST("searchGame", v1.SearchGame)
		gameRouter.POST("getGid", v1.GetGid)
	}
}
