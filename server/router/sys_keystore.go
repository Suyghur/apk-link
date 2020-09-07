/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_keystore
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitKeystoreRouter(r *gin.RouterGroup) {
	keystoreRouter := r.Group("keystore").Use(middleware.JWTAuth())
	{
		keystoreRouter.POST("getKeystores", v1.GetKeystores)
		keystoreRouter.POST("searchKeystore", v1.SearchKeystore)
		keystoreRouter.POST("create", v1.CreateKeystore)
	}
}
