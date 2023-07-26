package tc

import (
	"encoding/json"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsConf struct {
	CommonConf
	Appid      string
	AppKey     string
	SignName   string
	TemplateId string
}

type smsSvc struct {
	Conf *SmsConf
}

func NewSmsSvc(conf *SmsConf) *smsSvc {
	return &smsSvc{Conf: conf}
}
func (s *smsSvc) Send(phone string, temParams []*string) error {
	client, err := sms.NewClient(Credential(s.Conf.SecretId, s.Conf.SecretKey), regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		return err
	}

	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(s.Conf.Appid)
	req.SignName = common.StringPtr(s.Conf.SignName)
	req.TemplateId = common.StringPtr(s.Conf.TemplateId)
	req.TemplateParamSet = temParams
	req.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})

	resp, err := client.SendSms(req)
	if err != nil {
		logx.Errorf("[SMS] send failed. err:%s", err.Error())
		return err
	}

	b, _ := json.Marshal(resp)
	fmt.Printf("%s", string(b))

	return nil
}
