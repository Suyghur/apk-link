/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
)

func InitOriginRouter(r *gin.RouterGroup) {
	originRouter := r.Group("origin")
	{
		originRouter.POST("list", v1.ListOrigin)
		originRouter.POST("create", v1.CreateOrigin)
		originRouter.POST("modify", v1.ModifyOrigin)
	}

}
