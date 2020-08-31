/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : config
*/
package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Aliyun Aliyun `mapstructure:"aliyun" json:"aliyun" yaml:"aliyun"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use_multipoint" yaml:"use-multipoint"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	//DbType        string `mapstructure:"db-type" json:"db_type" yaml:"db-type"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max_idle_conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max_open_conns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"log_mode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Aliyun struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"access_key_id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access_key_secret" yaml:"access-key-secret"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region          string `mapstructure:"region" json:"region" yanl:"region"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	CName           bool   `mapstructure:"cname" json:"cname" yaml:"cname"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

//type Sqlite struct {
//}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signing_key" yaml:"signing-key"`
}

//type Casbin struct {
//}
