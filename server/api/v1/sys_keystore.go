/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_keystore
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

func SearchKeystore(c *gin.Context) {
	var bean request.ReqKeystoreListBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchKeystore(bean)
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

func GetKeystores(c *gin.Context) {
	var bean model.SysKeystore
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"GameSite": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GvaLog.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, keystores := service.GetKeystores(bean.GameSite)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(keystores, c)
	}
}

func CreateKeystore(c *gin.Context) {
	var bean model.SysKeystore
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameSite":              {utils.NotEmpty()},
		"KeystoreName":          {utils.NotEmpty()},
		"KeystorePassword":      {utils.NotEmpty()},
		"KeystoreAlias":         {utils.NotEmpty()},
		"KeystoreAliasPassword": {utils.NotEmpty()},
		"FingerprintsMD5":       {utils.NotEmpty()},
		"FingerprintsSHA1":      {utils.NotEmpty()},
		"FingerprintsSHA256":    {utils.NotEmpty()},
		"KeystoreFileUrl":       {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GvaLog.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.CreateKeystore(bean)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建签名文件成功", c)
	}
}

//func CreateKeystore(c *gin.Context) {
//	//TODO 生成签名
//	var bean request.ReqKeystoreBean
//	_ = c.ShouldBindJSON(&bean)
//	verifyRules := utils.Rules{
//		"GameGroup": {utils.NotEmpty()},
//	}
//	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
//		global.GvaLog.Error(verifyErr.Error())
//		response.FailWithMessage(verifyErr.Error(), c)
//		return
//	}
//	err := service.CreateKeystore(bean.GameSite)
//	if err != nil {
//		response.FailWithMessage(fmt.Sprintf("%v", err), c)
//	} else {
//		response.OkWithMessage("创建签名文件成功", c)
//	}
//}
