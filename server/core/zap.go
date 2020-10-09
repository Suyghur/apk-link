/*
@Time : 10/9/20
@Author : #Suyghur,
@File : zap
*/

package core

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/global"
	"server/utils"
	"time"
)

var (
	err    error
	level  zapcore.Level
	writer zapcore.WriteSyncer
)

func init() {
	//判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.GvaConfig.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GvaConfig.Zap.Director)
		_ = os.Mkdir(global.GvaConfig.Zap.Director, os.ModePerm)
	}

	//初始化配置文件Level
	switch global.GvaConfig.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel

	}
	//使用file-rotatelogs进行日志分割
	writer, err = getWriteSyncer()
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		global.GvaLog = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		global.GvaLog = zap.New(getEncoderCore())
	}
	if global.GvaConfig.Zap.ShowLine {
		global.GvaLog.WithOptions(zap.AddCaller())
	}
}

//getWriteSyncer zap logger中加入file-rotatelogs
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		global.GvaConfig.Zap.Director+string(os.PathSeparator)+"%Y-%m-%d.log",
		zaprotatelogs.WithLinkName(global.GvaConfig.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.GvaConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

//getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GvaConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.GvaConfig.Zap.EncodeLevel == "LowercaseLevelEncoder":
		//小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GvaConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		//小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GvaConfig.Zap.EncodeLevel == "CapitalLevelEncoder":
		//大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GvaConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		//大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return config
}

//getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.GvaConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

//getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() zapcore.Core {
	return zapcore.NewCore(getEncoder(), writer, level)
}

//自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(global.GvaConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
