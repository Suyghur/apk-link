/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_game
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/global/response"
	"server/model"
	"server/model/bean/request"
	resp "server/model/bean/response"
	"server/service"
	"server/utils"
)

func SearchChannel(c *gin.Context) {
	var bean request.ReqChannelListBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchChannel(bean)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     bean.Page,
			PageSize: bean.PageSize,
		}, c)
	}
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
