package wechatPay

import (
	"context"
	"fmt"
	"github.com/smallsha123/php2go/errorx"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"strconv"
	"time"
)

// PrepayWithRequestPayment Jsapi支付下单，并返回调起支付的请求参数
func (s *WechatPay) PrepayWithRequestPayment(
	ctx context.Context, req *jsapi.PrepayRequest,
) (resp *PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {
	s.Logger.Infof("【服务商模式微信JSAPI下单】请求参数：%s", req)

	if s.Client == nil {
		return nil, nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}

	// JSAPI下单
	prepayResp, result, err := s.PartnerJsapi(req)
	if err != nil {
		return nil, result, err
	}

	// 组装JSAPI调起支付参数
	resp = new(PrepayWithRequestPaymentResponse)
	resp.PrepayId = prepayResp.PrepayId
	resp.SignType = core.String("RSA")
	resp.Appid = req.SpAppid
	resp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		s.Logger.Errorf("【服务商模式微信JSAPI调起支付】随机字符串生成失败:%s", err.Error())
		return nil, nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "随机字符串生成失败")
	}
	resp.NonceStr = core.String(nonce)
	resp.Package = core.String("prepay_id=" + *prepayResp.PrepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *resp.Appid, *resp.TimeStamp, *resp.NonceStr, *resp.Package)
	signatureResult, err := s.Client.Sign(ctx, message)
	if err != nil {
		s.Logger.Errorf("【服务商模式微信JSAPI调起支付】签名失败:%s", err.Error())
		return nil, nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "签名失败")
	}
	resp.PaySign = core.String(signatureResult.Signature)
	return resp, result, nil
}

// PartnerJsapi 服务商模式微信支付
func (s WechatPay) PartnerJsapi(req *jsapi.PrepayRequest) (resp *jsapi.PrepayResponse, result *core.APIResult, err error) {

	if s.Client == nil {
		return nil, nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}

	svc := jsapi.JsapiApiService{Client: s.Client}

	resp, result, err = svc.Prepay(s.Ctx, *req)

	if err != nil {
		s.Logger.Errorf("【服务商模式微信JSAPI下单】错误信息:%s", err)
		return resp, result, s.getErr(err)
	}

	s.Logger.Infof("【服务商模式微信JSAPI下单】status=%d resp=%s", result.Response.StatusCode, resp)

	return resp, result, err
}

// CloseOrder 关闭订单
func (s WechatPay) CloseOrder(closeOrderRequest *jsapi.CloseOrderRequest) (err error) {
	svc := jsapi.JsapiApiService{Client: s.Client}
	_, err = svc.CloseOrder(s.Ctx, *closeOrderRequest)
	if err != nil {
		s.Logger.Errorf("【服务商模式微信关闭订单】错误信息:%s", err)
		return s.getErr(err)
	}
	return nil
}

// QueryOrder 查询订单
func (s WechatPay) QueryOrder(queryOrderByOutTradeNoRequest *jsapi.QueryOrderByOutTradeNoRequest) (resp *partnerpayments.Transaction, result *core.APIResult, err error) {
	if s.Client == nil {
		return nil, nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := jsapi.JsapiApiService{Client: s.Client}
	resp, result, err = svc.QueryOrderByOutTradeNo(s.Ctx, *queryOrderByOutTradeNoRequest)
	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式查询订单】call Prepay err:%s", err)
	} else {
		s.Logger.Infof("【服务商模式查询订单】status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return
}

// PartnerRefundDomestic 服务商模式退款
func (s WechatPay) PartnerRefundDomestic(req *PartnerRefundDomesticParams) (resp *refunddomestic.Refund, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := refunddomestic.RefundsApiService{Client: s.Client}
	resp, result, err := svc.Create(s.Ctx,
		refunddomestic.CreateRequest{
			SubMchid:      core.String(req.SubMchid),
			TransactionId: core.String(req.TransactionId),
			OutTradeNo:    core.String(req.OutTradeNo),
			OutRefundNo:   core.String(req.OutRefundNo),
			Reason:        core.String(req.Reason),
			NotifyUrl:     core.String(s.WechatPayParams.RefundNotify),
			FundsAccount:  refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(req.Refund),
				Total:    core.Int64(req.Total),
			},
		},
	)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式微信退款】请求数据：%+v， err:%s", req, err)
		return nil, s.getErr(err)
	}
	// 处理返回结果
	s.Logger.Infof("【服务商模式微信退款】status=%d resp=%s", result.Response.StatusCode, resp)

	return resp, nil
}
