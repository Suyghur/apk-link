/*
@Time : 2020/9/6
@Author : #Suyghur,
@File : sys_origins
*/

package response

type OriginsResponse struct {
	GameFileName    string `json:"game_file_name"`
	IsFuseSdk       int    `json:"is_fuse_sdk"`
	GameOrientation int    `json:"game_orientation"`
	ApkUrl          string `json:"apk_url"`
}
