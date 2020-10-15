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

func CreateGid(c *gin.Context) {
	var bean request.ReqGidBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"Gid":      {utils.NotEmpty()},
		"GameSite": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	game := &model.SysGid{
		Gid:      bean.Gid,
		GameSite: bean.GameSite,
	}
	err := service.CreateGid(*game)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建游戏ID成功", c)
	}
}

func GetGids(c *gin.Context) {
	var bean request.ReqGameBean
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"GameSite": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, gids := service.GetGids(bean.GameSite)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithData(gids, c)
	}
}

func SearchGame(c *gin.Context) {
	var bean request.ReqGameListBean
	_ = c.ShouldBindJSON(&bean)
	err, list, total := service.SearchGameSite(bean)
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

func GetGameSites(c*gin.Context){

}

func CreateGame(c *gin.Context) {
	//var bean request.ReqGameBean
	//_ = c.ShouldBindJSON(&bean)
	//verifyRules := utils.Rules{
	//	"GameSite": {utils.NotEmpty()},
	//}
	//if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
	//	response.FailWithMessage(verifyErr.Error(), c)
	//	return
	//}
	//game := &model.SysGame{
	//	GameSite: bean.GameSite,
	//}
	//err := service.CreateGame(*game)
	//if err != nil {
	//	global.GvaLog.Error(err.Error())
	//	response.FailWithMessage(fmt.Sprintf("%v", err), c)
	//} else {
	//	response.OkWithMessage("创建游戏成功", c)
	//}
}

func ModifyGame(c *gin.Context) {

}

func DelGame(c *gin.Context) {

}
