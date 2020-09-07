/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_base
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitBaseRouter(r *gin.RouterGroup) {
	baseRouter := r.Group("base")
	{
		baseRouter.POST("login", v1.Login)
		baseRouter.POST("register", v1.Register)
	}
}
