package wechatPay

import (
	"context"
	"crypto/rsa"
	"github.com/smallsha123/php2go/errorx"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/decryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"io/ioutil"
	"net/http"
)

type WechatPay struct {
	Client          *core.Client
	Ctx             context.Context
	MchPrivateKey   *rsa.PrivateKey
	WechatPayParams *Params
	Logger          logx.Logger
}

var ClientMap map[string]WechatPay

func InitWechatPay(params *Params) {

	ctx := context.Background()
	var (
		mchID                      = params.MchID                      // 商户号
		mchCertificateSerialNumber = params.MchCertificateSerialNumber // 商户证书序列号
		mchAPIv3Key                = params.MchAPIv3Key                // 商户APIv3密钥
		privateKey                 = params.PrivateKey                 // 私钥
	)

	// 从网络获取私钥
	PrivateKey := GetPrivateKey(privateKey)
	mchPrivateKey, err := utils.LoadPrivateKey(PrivateKey)
	if err != nil {
		logx.Errorf("【wechatPay实例化】load merchant private key error %+v", err)
		return
	}

	// 使用商户私钥等初始化 client
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
		option.WithWechatPayCipher(
			encryptors.NewWechatPayEncryptor(downloader.MgrInstance().GetCertificateVisitor(mchID)),
			decryptors.NewWechatPayDecryptor(mchPrivateKey),
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		logx.Errorf("【wechatPay实例化】new wechat minievent client err:%+v", err)
		return
	}
	ClientMap = make(map[string]WechatPay)
	ClientMap[params.MchID] = WechatPay{
		Client:          client,
		MchPrivateKey:   mchPrivateKey,
		WechatPayParams: params,
	}
}

// NewWechatPay 实例化
func NewWechatPay(ctx context.Context, params *Params) (*WechatPay, error) {

	logger := logx.WithContext(ctx)
	info, ok := ClientMap[params.MchID]
	if ok == false {
		logger.Errorf("【微信支付实例化】微信支付信息不存在：%+v", params)
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "微信支付信息不存在")
	}

	return &WechatPay{
		Client:          info.Client,
		Ctx:             ctx,
		MchPrivateKey:   info.MchPrivateKey,
		WechatPayParams: info.WechatPayParams,
		Logger:          logger,
	}, nil
}

// 获取错误信息
func (s WechatPay) getErr(errs error) (err error) {
	apiError, ok := errs.(*core.APIError)
	if ok {
		return errorx.NewCodeError(errorx.ERR_DEFAULT, apiError.Message)
	} else {
		return errs
	}

}

// GetPrivateKey 获取商户私钥信息
func GetPrivateKey(url string) string {
	v, err := http.Get(url)
	if err != nil {
		logx.Errorf("Http get [%v] failed! %v", url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logx.Errorf("关闭请求错误 %v", err)
		}
	}(v.Body)
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		logx.Errorf("Read http response failed! %v", err)
	}
	return string(content)
}
