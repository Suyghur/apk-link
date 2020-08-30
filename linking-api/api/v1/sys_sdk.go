/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"linking-api/global/response"
)

func ListFuseSdk(c *gin.Context) {
	response.OkWithData(gin.H{"list": "list"}, c)
}

func SearchFuseSdk(c *gin.Context) {
	version := c.Query("version")
	response.OkWithData(gin.H{
		"search":  "search",
		"version": version,
	}, c)
}

func ModifyFuseSdk(c *gin.Context) {
	response.OkWithData(gin.H{"list": "list"}, c)
}

func ListChannelSdk(c *gin.Context) {
	response.OkWithData(gin.H{"list": "list"}, c)
}

func SearchChannelSdk(c *gin.Context) {
	version := c.Query("version")
	response.OkWithData(gin.H{
		"search":  "search",
		"version": version,
	}, c)
}

func ListPluginSdk(c *gin.Context) {
	response.OkWithData(gin.H{"list": "list"}, c)
}

func SearchPluginSdk(c *gin.Context) {
	version := c.Query("version")
	response.OkWithData(gin.H{
		"search":  "search",
		"version": version,
	}, c)
}
