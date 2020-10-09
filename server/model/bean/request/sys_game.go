/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

type ReqGameBean struct {
	GameGroup string `json:"game_group" form:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid       string `json:"gid" form:"gid" gorm:"not null;index:idx_gid;comment:'游戏组ID'"`
}

type ReqGameListBean struct {
	Gid string `json:"gid" form:"gid" gorm:"not null;index:idx_gid;comment:'游戏组ID'"`
	PageInfo
}
