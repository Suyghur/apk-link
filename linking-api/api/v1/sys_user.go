/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : sys_user
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"linking-api/global/response"
)

func Test(c *gin.Context) {
	response.OkWithData(gin.H{"test": "test"}, c)
}
