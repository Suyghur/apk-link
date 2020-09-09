/*
@Time : 2020/9/8
@Author : #Suyghur,
@File : sys_link
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitLinkRouter(r *gin.RouterGroup) {
	linkRouter := r.Group("link")
	{
		linkRouter.POST("searchLink", v1.SearchLink)
	}
}
