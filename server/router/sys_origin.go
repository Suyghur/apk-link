/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitOriginRouter(r *gin.RouterGroup) {
	originRouter := r.Group("origin").Use(middleware.JWTAuth())
	{
		originRouter.GET("getOrigins", v1.GetOrigins)
		originRouter.GET("searchOrigin", v1.SearchOrigin)
		originRouter.POST("createOrigin", v1.CreateOrigin)
		originRouter.PUT("modifyOrigin", v1.ModifyOrigin)
	}
}
