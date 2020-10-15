/*
@Time : 2020/9/8
@Author : #Suyghur,
@File : sys_link
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

func ReportLink(c *gin.Context) {
	var bean request.ReqLinkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"TaskId":         {utils.NotEmpty()},
		"GameSite":       {utils.NotEmpty()},
		"GameName":       {utils.NotEmpty()},
		"Aid":            {utils.NotEmpty()},
		"FuseSdkVersion": {utils.NotEmpty()},
		"ChannelName":    {utils.NotEmpty()},
		"PluginName":     {utils.NotEmpty()},
		"LinkUrl":        {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GvaLog.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}

	link := &model.SysLink{
		TaskId:         bean.TaskId,
		GameSite:       bean.GameSite,
		GameName:       bean.GameName,
		Aid:            bean.Aid,
		FuseSdkVersion: bean.FuseSdkVersion,
		ChannelName:    bean.ChannelName,
		PluginName:     bean.PluginName,
		LinkUrl:        bean.LinkUrl,
	}

	err := service.ReportLink(*link)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("上报Link成功", c)
	}
}

func SearchLink(c *gin.Context) {
	var bean request.ReqLinkListBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchLink(bean)
	if err != nil {
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
