/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_options
*/

package response

type OptionsResponse struct {
	GameOptions    []string `json:"game_options"`
	ChannelOptions []string `json:"channel_options"`
	PluginOptions  []string `json:"plugin_options"`
}
