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
	"server/model/bean/request"
	resp "server/model/bean/response"
	"server/service"
	"server/utils"
)

func SearchKeystore(c *gin.Context) {
	var bean request.ReqKeystoreListBean
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
}

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
