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

func ListGame(c *gin.Context) {
	//var bean request.ReqListFuseSdkBean
	//_ = c.ShouldBindJSON(&bean)
	err, result := service.ListGame()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}

}

func CreateGame(c *gin.Context) {
	var bean request.ReqGameBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup": {utils.NotEmpty()},
		"Gid":       {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	game := &model.SysGame{
		GameGroup: bean.GameGroup,
		Gid:       bean.Gid,
	}
	err := service.CreateGame(*game)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建游戏成功", c)
	}
}

func ModifyGame(c *gin.Context) {

}

func DelGame(c *gin.Context) {

}

func GetGid(c *gin.Context) {
	var bean request.ReqGameBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	//game := &model.SysGame{
	//	GameGroup: bean.GameGroup,
	//	Gid:       bean.Gid,
	//}
	err, game := service.GetGid(bean.GameGroup)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithData(gin.H{"gid": game.Gid}, c)
	}
}
