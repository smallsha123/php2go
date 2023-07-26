package wechatPay

import (
	"github.com/smallsha123/php2go/errorx"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

// ProfitSharingCreateOrder 请求分账API
func (s *WechatPay) ProfitSharingCreateOrder(request profitsharing.CreateOrderRequest) (resp *profitsharing.OrdersEntity, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	s.Logger.Infof("【服务商模式请求分账】请求参数：%s", request)
	svc := profitsharing.OrdersApiService{Client: s.Client}
	resp, result, err := svc.CreateOrder(s.Ctx, request)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式请求分账】 err:%s", err.Error())
		return nil, s.getErr(err)

	}
	s.Logger.Infof("status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

// ProfitSharingQueryOrder 查询分账结果
func (s *WechatPay) ProfitSharingQueryOrder(queryOrderRequest profitsharing.QueryOrderRequest) (resp *profitsharing.OrdersEntity, err error) {

	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.OrdersApiService{Client: s.Client}
	resp, result, err := svc.QueryOrder(s.Ctx, queryOrderRequest)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式查询分账结果】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式查询分账结果】status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

// ProfitSharingReturnOrder 请求分账回退
func (s *WechatPay) ProfitSharingReturnOrder(createReturnOrder profitsharing.CreateReturnOrderRequest) (resp *profitsharing.ReturnOrdersEntity, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.ReturnOrdersApiService{Client: s.Client}
	resp, result, err := svc.CreateReturnOrder(s.Ctx, createReturnOrder)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式请求分账回退】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式请求分账回退】status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

// ProfitSharingUnfreezeOrder 解冻剩余资金
func (s *WechatPay) ProfitSharingUnfreezeOrder(createReturnOrder profitsharing.UnfreezeOrderRequest) (resp *profitsharing.OrdersEntity, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.OrdersApiService{Client: s.Client}
	resp, result, err := svc.UnfreezeOrder(s.Ctx, createReturnOrder)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式解冻剩余资金】err:%s", err.Error())
		return nil, s.getErr(err)

	}
	s.Logger.Infof("【服务商模式解冻剩余资金】status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

// ProfitSharingQuerySurplusAmounts 查询剩余待分金额API
func (s *WechatPay) ProfitSharingQuerySurplusAmounts(QueryOrderAmountRequest profitsharing.QueryOrderAmountRequest) (resp *profitsharing.QueryOrderAmountResponse, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.TransactionsApiService{Client: s.Client}
	resp, result, err := svc.QueryOrderAmount(s.Ctx, QueryOrderAmountRequest)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式查询剩余待分金额】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式查询剩余待分金额】status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

// ProfitSharingReceiversAdd 添加分账接收方
func (s *WechatPay) ProfitSharingReceiversAdd(addReceiverRequest profitsharing.AddReceiverRequest) (resp *profitsharing.AddReceiverResponse, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.ReceiversApiService{Client: s.Client}
	resp, result, err := svc.AddReceiver(s.Ctx, addReceiverRequest)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式添加分账接收方】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式添加分账接收方】status=%d resp=%s", result.Response.Body, resp)
	return resp, nil
}

// ProfitSharingReceiversDelete 删除分账接收方
func (s *WechatPay) ProfitSharingReceiversDelete(deleteReceiverRequest profitsharing.DeleteReceiverRequest) (resp *profitsharing.DeleteReceiverResponse, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.ReceiversApiService{Client: s.Client}
	resp, result, err := svc.DeleteReceiver(s.Ctx, deleteReceiverRequest)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式删除分账接收方】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式删除分账接收方】status=%d resp=%s", result.Response.Body, resp)
	return resp, nil
}

// ProfitSharingMerchantConfig 查询最大分账比例
func (s *WechatPay) ProfitSharingMerchantConfig(queryMerchantRatioRequest profitsharing.QueryMerchantRatioRequest) (resp *profitsharing.QueryMerchantRatioResponse, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	svc := profitsharing.MerchantsApiService{Client: s.Client}
	resp, result, err := svc.QueryMerchantRatio(s.Ctx, queryMerchantRatioRequest)

	if err != nil {
		// 处理错误
		s.Logger.Errorf("【服务商模式查询最大分账比例】err:%s", err.Error())
		return nil, s.getErr(err)
	}
	s.Logger.Infof("【服务商模式查询最大分账比例】status=%d resp=%s", result.Response.Body, resp)
	return resp, nil
}
