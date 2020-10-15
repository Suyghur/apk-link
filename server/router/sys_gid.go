/*
@Time : 10/10/20
@Author : #Suyghur,
@File : sys_gid
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitGidRouter(r *gin.RouterGroup) {
	gidRouter := r.Group("gid").Use(middleware.JWTAuth())
	{
		gidRouter.POST("createGid", v1.CreateGid)
		gidRouter.GET("getGids", v1.GetGids)
	}
}
