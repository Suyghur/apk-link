/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
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

func GetOrigins(c *gin.Context) {
	var bean request.ReqGetOriginListBean
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"GameGroup": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GVA_LOG.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, origins := service.GetOrigins(bean.GameGroup)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"origins": origins}, c)
	}
}

func SearchOrigin(c *gin.Context) {
	var bean request.ReqSearchOriginBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchOrigin(bean)
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

func CreateOrigin(c *gin.Context) {
	var bean request.ReqOriginBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":       {utils.NotEmpty()},
		"Gid":             {utils.NotEmpty()},
		"IsFuseSdk":       {utils.NotEmpty()},
		"SdkVersion":      {utils.NotEmpty()},
		"GameFileName":    {utils.NotEmpty()},
		"GameVersionCode": {utils.NotEmpty()},
		"GameVersionName": {utils.NotEmpty()},
		"KeystoreName":    {utils.NotEmpty()},
		"ApkUrl":          {utils.NotEmpty()},
	}

	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	origin := &model.SysOrigin{
		GameGroup:       bean.GameGroup,
		Gid:             bean.Gid,
		IsFuseSdk:       bean.IsFuseSdk,
		SdkVersion:      bean.SdkVersion,
		GameFileName:    bean.GameFileName,
		GameVersionCode: bean.GameVersionCode,
		GameVersionName: bean.GameVersionName,
		GameOrientation: bean.GameOrientation,
		KeystoreName:    bean.KeystoreName,
		IconUrl:         bean.IconUrl,
		ApkUrl:          bean.ApkUrl,
	}
	err := service.CreateOrigin(*origin)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建游戏母包成功", c)
	}
}

func ModifyOrigin(c *gin.Context) {
	var bean request.ReqOriginBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":       {utils.NotEmpty()},
		"Gid":             {utils.NotEmpty()},
		"IsFuseSdk":       {utils.NotEmpty()},
		"SdkVersion":      {utils.NotEmpty()},
		"GameVersionCode": {utils.NotEmpty()},
		"GameVersionName": {utils.NotEmpty()},
		"KeystoreName":    {utils.NotEmpty()},
		"ApkUrl":          {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	origin := &model.SysOrigin{
		GameGroup:       bean.GameGroup,
		Gid:             bean.Gid,
		IsFuseSdk:       bean.IsFuseSdk,
		SdkVersion:      bean.SdkVersion,
		GameVersionCode: bean.GameVersionCode,
		GameVersionName: bean.GameVersionName,
		GameOrientation: bean.GameOrientation,
		KeystoreName:    bean.KeystoreName,
		IconUrl:         bean.IconUrl,
		ApkUrl:          bean.ApkUrl,
	}
	err := service.ModifyOrigin(*origin)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改游戏母包成功", c)
	}
}
