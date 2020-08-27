/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : cros
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//处理跨域请求，支持options访问
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		//放行所有options方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}

		//处理请求
		context.Next()
	}
}
