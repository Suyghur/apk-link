/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package request

import uuid "github.com/satori/go.uuid"

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid"`
}
