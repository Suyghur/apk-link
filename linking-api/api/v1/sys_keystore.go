/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_keystore
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linking-api/global"
	"linking-api/global/response"
	"linking-api/model/bean/request"
	resp "linking-api/model/bean/response"
	"linking-api/service"
	"linking-api/utils"
)

func SearchKeystore(c *gin.Context) {
	var bean request.ReqListKeystoreBean
	_ = c.ShouldBindJSON(&bean)
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
	//var pageInfo request.PageInfo
	//_ = c.ShouldBindJSON(&pageInfo)
	//PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	//if PageVerifyErr != nil {
	//	response.FailWithMessage(PageVerifyErr.Error(), c)
	//	return
	//}
	//err, list, total := service.SearchKeystore(pageInfo)
}

//func SearchKeystore(c *gin.Context) {
//
//}

//func ModifyKeystore(c *gin.Context) {
//	//var reqStruct request.KeystoreStruct
//	//_ := c.ShouldBindJSON(&reqStruct)
//	//keystoreVerify := utils.Rules{
//	//	"keystore_name":           {utils.NotEmpty()},
//	//	"keystore_password":       {utils.NotEmpty()},
//	//	"keystore_alias":          {utils.NotEmpty()},
//	//	"keystore_alias_password": {utils.NotEmpty()},
//	//}
//	//keystoreVerifyErr := utils.Verify(reqStruct, keystoreVerify)
//	//if keystoreVerifyErr != nil {
//	//	response.FailWithMessage(keystoreVerifyErr.Error(), c)
//	//}
//
//}

func GetKeystores(c *gin.Context) {
	var bean request.ReqKeystoreBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GVA_LOG.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, keystores := service.GetKeystores(bean.GameGroup)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"keystores": keystores}, c)
	}
}

func CreateKeystore(c *gin.Context) {
	//TODO 生成签名
	var bean request.ReqKeystoreBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		global.GVA_LOG.Error(verifyErr.Error())
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.CreateKeystore(bean.GameGroup)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建签名文件成功", c)
	}
}
