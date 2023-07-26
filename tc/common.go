package tc

import "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

type CommonConf struct {
	SecretId  string
	SecretKey string
}

func Credential(secretId, secretKey string) *common.Credential {
	return common.NewCredential(secretId, secretKey)
}
