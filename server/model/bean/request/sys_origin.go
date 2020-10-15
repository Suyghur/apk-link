/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package request

import "server/model"

type ReqSearchOriginBean struct {
	model.SysOrigin
	PageInfo
}
