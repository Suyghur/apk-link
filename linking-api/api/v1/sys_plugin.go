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

func ListPlugin(c *gin.Context) {

}

func CreatePlugin(c *gin.Context) {
	var bean request.ReqPluginBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"PluginName":  {utils.NotEmpty()},
		"PluginAlias": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	plugin := &model.SysPlugin{
		PluginName:  bean.PluginName,
		PluginAlias: bean.PluginAlias,
	}
	err := service.CreatePlugin(*plugin)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建插件成功", c)
	}
}

func ModifyPlugin(c *gin.Context) {

}

func DelPlugin(c *gin.Context) {

}
