/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_options
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/global/response"
	"server/service"
)

func GetOptions(c *gin.Context) {
	err, options := service.GetOptions()
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"options": options}, c)
	}
}
