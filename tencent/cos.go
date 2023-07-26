package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
	"os"
)

type Cos struct {
	client *cos.Client
	config CosConfig
}

func NewCosClient(config CosConfig) *Cos {
	urlStr, _ := url.Parse("https://" + config.Bucket + ".cos." + config.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	return &Cos{
		client: cos.NewClient(baseURL, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.SecretId,
				SecretKey: config.SecretKey,
			},
		}),
		config: config,
	}
}

func (l *Cos) GetBaseUrl() string {
	return "https://" + l.config.Bucket + ".cos." + l.config.Region + ".myqcloud.com"
}

func (l *Cos) UploadFile(localFilePath, saveFilePath string) (string, error) {
	f, err := os.OpenFile(localFilePath, os.O_RDONLY, 0660)
	if err != nil {
		logx.Errorf("function file.Open() Filed err : %s", err)
	}
	defer f.Close()

	_, err = l.client.Object.Put(context.Background(), saveFilePath, f, nil)
	if err != nil {
		return "", err
	}

	return l.GetBaseUrl() + "/" + saveFilePath, nil
}
