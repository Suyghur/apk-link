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
	"server/model/bean/request"
	"server/service"
	"server/utils"
)

func GetOptions(c *gin.Context) {
	err, options := service.GetOptions()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"options": options}, c)
	}
}

func GetAids(c *gin.Context) {
	var bean request.ReqAidsBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"Gid": {utils.NotEmpty()},
	}

	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, aids := service.GetAids(bean.Gid)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"aids": aids}, c)
	}
}
