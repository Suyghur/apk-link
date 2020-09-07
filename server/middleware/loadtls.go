/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : loadtls
*/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

//用https把这个中间件在router里面use一下
func LoadTls() gin.HandlerFunc {
	return func(context *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(context.Writer, context.Request)
		if err != nil {
			//如果出现错误，不要继续
			fmt.Println(err)
			return
		}
		context.Next()
	}
}
