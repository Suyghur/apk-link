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

func SearchPlugin(c *gin.Context) {
	var bean request.ReqPluginListBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchPlugin(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
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
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建插件成功", c)
	}
}

func ModifyPlugin(c *gin.Context) {

}

func DeletePlugin(c *gin.Context) {

}
