/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
	"linking-api/middleware"
)

func InitOriginRouter(r *gin.RouterGroup) {
	originRouter := r.Group("origin").Use(middleware.JWTAuth())
	{
		originRouter.POST("getOrigins", v1.GetOrigins)
		originRouter.POST("searchOrigin", v1.SearchOrigin)
		originRouter.POST("createOrigin", v1.CreateOrigin)
		originRouter.POST("modifyOrigin", v1.ModifyOrigin)
	}

}
