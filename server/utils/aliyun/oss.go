/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : oss
*/

package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"server/global"
)

func OssSdkVersion() {
	global.GvaLog.Debug("aliyun oss sdk version: " + oss.Version)
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		global.GvaLog.Info("开始传输...", zap.Int64("consumed_bytes", event.ConsumedBytes), zap.Int64("consumed_bytes", event.TotalBytes))
	case oss.TransferDataEvent:
		global.GvaLog.Info("正在传输...",
			zap.Int64("consumed_bytes", event.ConsumedBytes),
			zap.Int64("total_bytes", event.TotalBytes),
			zap.Int64("percent", event.ConsumedBytes*100/event.TotalBytes))
	case oss.TransferCompletedEvent:
		global.GvaLog.Info("传输完成...",
			zap.Int64("consumed_bytes", event.ConsumedBytes),
			zap.Int64("total_bytes", event.TotalBytes))
	case oss.TransferFailedEvent:
		global.GvaLog.Error("传输失败...",
			zap.Int64("consumed_bytes", event.ConsumedBytes),
			zap.Int64("total_bytes", event.TotalBytes))
	default:
	}
}

//小文件上传
func Upload(targetBucketPath string, localFilePath string) {
	client, err := oss.New(global.GvaConfig.Aliyun.Endpoint, global.GvaConfig.Aliyun.AccessKeyId,
		global.GvaConfig.Aliyun.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		global.GvaLog.Error(err.Error())
		return
	}
	bucket, err := client.Bucket(global.GvaConfig.Aliyun.Bucket)
	if err != nil {
		global.GvaLog.Error(err.Error())
		return
	}

	if err := bucket.PutObjectFromFile(targetBucketPath, localFilePath, oss.Progress(&OssProgressListener{})); err != nil {
		global.GvaLog.Error(err.Error())
		return
	} else {
		global.GvaLog.Info(localFilePath + "上传成功")
	}
}

//小文件下载
func Download(fileUrl, saveFilePath string) {
	client, err := oss.New(global.GvaConfig.Aliyun.Endpoint, global.GvaConfig.Aliyun.AccessKeyId,
		global.GvaConfig.Aliyun.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		global.GvaLog.Error(err.Error())
		return
	}
	bucket, err := client.Bucket(global.GvaConfig.Aliyun.Bucket)
	if err != nil {
		global.GvaLog.Error(err.Error())
		return
	}
	if err := bucket.GetObjectToFileWithURL(fileUrl, saveFilePath, oss.Progress(&OssProgressListener{})); err != nil {
		global.GvaLog.Error(err.Error())
		return
	} else {
		global.GvaLog.Info(saveFilePath + "下载成功")
	}
}
