/*
@Time : 2020/9/1
@Author : #Suyghur,
@File : sys_keystore
*/

package response

import "linking-api/model"

type SysKeystoreResponse struct {
	Keystore model.SysKeystore `json:"keystore"`
}
