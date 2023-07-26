package elasticsearch

import (
	json2 "encoding/json"
	elasticsearch2 "github.com/elastic/go-elasticsearch/v7"
	"github.com/smallsha123/php2go/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type EsConfig struct {
	Host     string
	UserName string
	PassWord string
}

var Config *EsConfig

func InitEsConfig(esConfig *EsConfig) {
	Config = &EsConfig{
		Host:     esConfig.Host,
		UserName: esConfig.UserName,
		PassWord: esConfig.PassWord,
	}
}
func GetEsClient() (*elasticsearch2.Client, error) {
	esConfig := elasticsearch2.Config{
		Addresses: []string{
			Config.Host,
		},
		Username: Config.UserName,
		Password: Config.PassWord,
	}
	es, err := elasticsearch2.NewClient(esConfig)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error creating the client:"+err.Error())
	}
	_, err = es.Info()
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error creating the client:"+err.Error())
	}
	return es, err
}

func EsBodyLog(body map[string]interface{}) {
	str, _ := json2.Marshal(body)
	logx.Infof("Es body: %s ", str)
}
