/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

type ReqChannelBean struct {
	ChannelName  string `json:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	ChannelAlias string `json:"channel_alias" gorm:"not null;index:idx_channel_alias;comment:'渠道别名'"`
}
