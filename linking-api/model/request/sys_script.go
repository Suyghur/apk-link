/*
@Time : 9/2/2020
@Author : #Suyghur,
@File : sys_script
*/

package request

type CreateScriptStruct struct {
	GameGroup      string `json:"game_group" gorm:"not null;comment:'游戏组'"`
	ScriptFileName string `json:"script_file_name" gorm:"not null;comment:'脚本文件名'"`
	ScriptFileUrl  string `json:"script_file_url" gorm:"not null;comment:'脚本文件链接'"`
	ScriptFileMD5  string `json:"script_file_md5" gorm:"not null;comment:'脚本文件MD5值'"`
}
