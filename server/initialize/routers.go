/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : routers
*/
package initialize

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	//开启https
	//router.User(middleware.LoadTls())

	//跨域
	router.Use(middleware.Cors())
	return router
}
