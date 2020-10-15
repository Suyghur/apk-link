/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

import "server/model"

type ReqGameListBean struct {
	model.SysGame
	PageInfo
}
