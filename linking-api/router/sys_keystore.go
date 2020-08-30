/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_keystore
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
)

func InitKeystoreRouter(r *gin.RouterGroup) {
	keystoreRouter := r.Group("keystore")
	{
		keystoreRouter.POST("list", v1.ListKeystore)
		keystoreRouter.POST("search", v1.SearchKeystore)
		keystoreRouter.POST("modify", v1.ModifyKeystore)
	}
}
