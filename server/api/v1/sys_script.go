/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_script
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

func SearchScript(c *gin.Context) {
	var bean request.ReqListScriptBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchScript(bean)
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

func ModifyScript(c *gin.Context) {
	var bean request.ReqScriptBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":      {utils.NotEmpty()},
		"ScriptFileName": {utils.NotEmpty()},
		"ScriptFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	script := &model.SysScript{
		GameGroup:      bean.GameGroup,
		ScriptFileName: bean.ScriptFileName,
		ScriptFileUrl:  bean.ScriptFileUrl,
	}
	err := service.ModifyScript(*script)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改游戏脚本成功", c)
	}
}

func CreateScript(c *gin.Context) {
	var bean request.ReqScriptBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":      {utils.NotEmpty()},
		"ScriptFileName": {utils.NotEmpty()},
		"ScriptFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	script := &model.SysScript{
		GameGroup:      bean.GameGroup,
		ScriptFileName: bean.ScriptFileName,
		ScriptFileUrl:  bean.ScriptFileUrl,
	}
	err := service.CreateScript(*script)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建游戏脚本成功", c)
	}
}
