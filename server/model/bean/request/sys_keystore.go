/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_keystore
*/

package request

import "server/model"

type ReqKeystoreListBean struct {
	model.SysKeystore
	PageInfo
}
