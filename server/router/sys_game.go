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
		gameRouter.DELETE("deleteGame", v1.DelGame)
		gameRouter.PUT("modifyGame", v1.ModifyGame)
		gameRouter.GET("searchGame", v1.SearchGame)
		gameRouter.GET("getGid", v1.GetGid)
	}
}
