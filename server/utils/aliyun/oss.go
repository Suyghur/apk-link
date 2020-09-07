/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : oss
*/

package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"server/global"
)

func OssSdkVersion() {
	global.GVA_LOG.Debug("aliyun oss sdk version: ", oss.Version)
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		global.GVA_LOG.Infof("开始传输... ConsumedBytes: %d, TotalBytes %d",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		global.GVA_LOG.Infof("正在传输... ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		global.GVA_LOG.Infof("传输完成... ConsumedBytes: %d, TotalBytes %d.",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		global.GVA_LOG.Errorf("传输失败, ConsumedBytes: %d, TotalBytes %d.",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

//小文件上传
func Upload(targetBucketPath string, localFilePath string) {
	client, err := oss.New(global.GVA_CONFIG.Aliyun.Endpoint, global.GVA_CONFIG.Aliyun.AccessKeyId,
		global.GVA_CONFIG.Aliyun.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	bucket, err := client.Bucket(global.GVA_CONFIG.Aliyun.Bucket)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	if err := bucket.PutObjectFromFile(targetBucketPath, localFilePath, oss.Progress(&OssProgressListener{})); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	} else {
		global.GVA_LOG.Info(localFilePath + "上传成功")
	}
}

//小文件下载
func Download(fileUrl, saveFilePath string) {
	client, err := oss.New(global.GVA_CONFIG.Aliyun.Endpoint, global.GVA_CONFIG.Aliyun.AccessKeyId,
		global.GVA_CONFIG.Aliyun.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	bucket, err := client.Bucket(global.GVA_CONFIG.Aliyun.Bucket)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	if err := bucket.GetObjectToFileWithURL(fileUrl, saveFilePath, oss.Progress(&OssProgressListener{})); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	} else {
		global.GVA_LOG.Info(saveFilePath + "下载成功")
	}
}
