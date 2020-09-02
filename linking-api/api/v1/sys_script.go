/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_script
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linking-api/global/response"
	"linking-api/model"
	"linking-api/model/request"
	resp "linking-api/model/response"
	"linking-api/service"
	"linking-api/utils"
)

func ListGameScript(c *gin.Context) {

}

func SearchGameScript(c *gin.Context) {

}

func ModifyGameScript(c *gin.Context) {

}

func CreateGameScript(c *gin.Context) {
	var reqStruct request.CreateScriptStruct
	_ := c.ShouldBindJSON(&reqStruct)
	verifyRules := utils.Rules{
		"game_group":       {utils.NotEmpty()},
		"script_file_name": {utils.NotEmpty()},
		"script_file_url":  {utils.NotEmpty()},
		"script_file_md5":  {utils.NotEmpty()},
	}
	verifyErr := utils.Verify(reqStruct, verifyRules)
	if verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}

	script := &model.SysScript{
		GameGroup:      reqStruct.GameGroup,
		ScriptFileName: reqStruct.ScriptFileName,
		ScriptFileUrl:  reqStruct.ScriptFileUrl,
		ScriptFileMD5:  reqStruct.ScriptFileMD5,
	}
	err, scriptReturn := service.CreateScript(*script)
	if err !=nil{
		response.FailWithDetail(response.ERROR, fmt.Sprintf("%v", err), resp.SysUserResponse{User: userReturn}, c)

	}
}
