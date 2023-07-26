package tc

import (
	"github.com/smallsha123/php2go/errorx"
	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	captchaType         uint64 = 9
	CaptchaAppid        uint64 = 2012878865
	ipAddr                     = "127.0.0.1"
	CaptchaAppSecretKey        = "0dJgDxJ3INCrZ8H1KKcLZqw**"
)

type CaptchaConf struct {
	CommonConf
	// AppId     uint64
	// AppSecret string
}

type captchaSvc struct {
	Conf *CaptchaConf
}

func NewCaptchaSvc(conf *CaptchaConf) *captchaSvc {
	return &captchaSvc{
		Conf: conf,
	}
}

func (s *captchaSvc) VerifyCaptcha(ticket, randstr string) error {
	client, err := captcha.NewClient(Credential(s.Conf.SecretId, s.Conf.SecretKey), regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		logx.Errorf("[VerifyCaptcha] init captcha client failed. err:%s", err.Error())
		return err
	}

	req := captcha.NewDescribeCaptchaResultRequest()
	req.Ticket = &ticket
	req.Randstr = &randstr
	req.AppSecretKey = &CaptchaAppSecretKey
	req.CaptchaAppId = &CaptchaAppid
	req.CaptchaType = &captchaType
	req.UserIp = &ipAddr
	resp, err := client.DescribeCaptchaResult(req)
	if err != nil {
		return err
	}

	if *resp.Response.CaptchaCode != 1 {
		logx.Errorf("[VerifyCaptcha] response failed. err:%s", *resp.Response.CaptchaMsg)
		return errorx.NewCodeError(errorx.ERR_DEFAULT, *resp.Response.CaptchaMsg)
	}

	return nil
}
