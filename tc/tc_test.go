package tc

import "testing"

func TestVerifyCaptcha(t *testing.T) {
	err := NewCaptchaSvc(&CaptchaConf{CommonConf{
		SecretId:  "",
		SecretKey: "",
	}}).VerifyCaptcha("xxx", "xxx")
	t.Log(err)
}
