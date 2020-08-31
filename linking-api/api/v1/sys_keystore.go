/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_keystore
*/

package v1

import (
	"github.com/gin-gonic/gin"
)

func ListKeystore(c *gin.Context) {
	//var pageInfo request.PageInfo
	//_ = c.ShouldBindJSON(&pageInfo)
	//PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	//if PageVerifyErr != nil {
	//	response.FailWithMessage(PageVerifyErr.Error(), c)
	//	return
	//}
	//err, list, total := service.ListKeystore(pageInfo)
}

func SearchKeystore(c *gin.Context) {

}

func ModifyKeystore(c *gin.Context) {
	//var reqStruct request.KeystoreStruct
	//_ := c.ShouldBindJSON(&reqStruct)
	//keystoreVerify := utils.Rules{
	//	"keystore_name":           {utils.NotEmpty()},
	//	"keystore_password":       {utils.NotEmpty()},
	//	"keystore_alias":          {utils.NotEmpty()},
	//	"keystore_alias_password": {utils.NotEmpty()},
	//}
	//keystoreVerifyErr := utils.Verify(reqStruct, keystoreVerify)
	//if keystoreVerifyErr != nil {
	//	response.FailWithMessage(keystoreVerifyErr.Error(), c)
	//}

}

func CreateKeystore(c *gin.Context) {
	//TODO 生成签名
	//var
}
