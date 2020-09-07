/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : response
*/

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR   = 1
	SUCCESS = 0

	ERROR_MSG   = "error"
	SUCCESS_MSG = "success"
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, SUCCESS_MSG, map[string]interface{}{}, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, msg, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, SUCCESS_MSG, data, c)
}

func OkWithDetailed(msg string, data interface{}, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, ERROR_MSG, map[string]interface{}{}, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, msg, map[string]interface{}{}, c)
}

func FailWithDetail(code int, msg string, data interface{}, c *gin.Context) {
	Result(code, msg, data, c)
}
