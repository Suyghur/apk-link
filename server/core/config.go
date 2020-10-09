/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : config
*/

package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "server/packfile"

	"server/global"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		fmt.Println(err)
	}
	global.GvaVp = v
}
