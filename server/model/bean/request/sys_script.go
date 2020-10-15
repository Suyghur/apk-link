/*
@Time : 9/2/2020
@Author : #Suyghur,
@File : sys_script
*/

package request

type ReqScriptBean struct {
	GameSite       string `json:"game_site" form:"game_site" gorm:"not null;comment:'游戏组标识'"`
	ScriptFileName string `json:"script_file_name" form:"script_file_name" gorm:"not null;comment:'脚本文件名'"`
	ScriptFileUrl  string `json:"script_file_url" form:"script_file_url" gorm:"not null;comment:'脚本文件链接'"`
}

type ReqListScriptBean struct {
	ReqScriptBean
	PageInfo
}
