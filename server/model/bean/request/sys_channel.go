/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

type ReqChannelBean struct {
	Cid          string `json:"cid" form:"cid" gorm:"not null;index:idx_cid;comment:'渠道ID'"`
	ChannelName  string `json:"channel_name" form:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	ChannelAlias string `json:"channel_alias" form:"channel_alias" gorm:"not null;index:idx_channel_alias;comment:'渠道别名'"`
}

type ReqChannelListBean struct {
	ChannelAlias string `json:"channel_alias" form:"channel_alias" gorm:"not null;index:idx_channel_alias;comment:'渠道别名'"`
	PageInfo
}
