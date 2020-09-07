/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_game
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linking-api/global"
	"linking-api/global/response"
	"linking-api/model"
	"linking-api/model/bean/request"
	"linking-api/service"
	"linking-api/utils"
)

func ListChannel(c *gin.Context) {

}

func CreateChannel(c *gin.Context) {
	var bean request.ReqChannelBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"ChannelName":  {utils.NotEmpty()},
		"ChannelAlias": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	channel := &model.SysChannel{
		ChannelName:  bean.ChannelName,
		ChannelAlias: bean.ChannelAlias,
	}
	err := service.CreateChannel(*channel)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建渠道成功", c)
	}
}

func ModifyChannel(c *gin.Context) {

}

func DelChannel(c *gin.Context) {

}
