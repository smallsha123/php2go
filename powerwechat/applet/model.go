package applet

import "github.com/zeromicro/go-zero/core/stores/redis"

// PowerWechatParams 基础参数
type PowerWechatParams struct {
	AppID           string // 小程序appid
	Secret          string // 小程序app secret
	AppName         string // 小程序名称
	Token           string // 小程序Token
	AESKey          string // 小程序AESKey
	HttpDebug       bool
	Level           string
	File            string
	Redis           *redis.Redis `json:"redis,optional"`
	PayTemplate     string
	WaitPayTemplate string
	ExpressTemplate string
}

// RegisterCustomerServiceReq 注册商户子客服请求
type RegisterCustomerServiceReq struct {
	AccountName string `json:"account_name"`
	Nickname    string `json:"nickname"`
	IconMediaID string `json:"icon_media_id"`
}

// RegisterCustomerServiceReply 注册商户子客服返回
type RegisterCustomerServiceReply struct {
	BusinessId int64 `json:"business_id"`
}

// RegisterMiniServiceReq  回复小程序投诉
type RegisterMiniServiceReq struct {
	Content          string   `json:"content"`          // 投诉内容
	ComplaintOrderId string   `json:"complaintOrderId"` // 投诉ID
	MediaIdList      []string `json:"mediaIdList"`      // 图片素材
	BussiHandle      int64    `json:"bussiHandle"`      // 操作1是同意和解，2是拒绝和解
}

// WxInfoReply 获取投诉详情
type WxInfoReply struct {
	Errcode        int64              `json:"errcode"` // 投诉内容
	Errmsg         string             `json:"errmsg"`  // 投诉ID
	ComplaintOrder ComplaintOrderInfo `json:"complaintOrder"`
	Item           []Present          `json:"item"` // 投诉进度

}

// present   投诉进度
type Present struct {
	ItemType    int64    `json:"itemType"`    //投诉节点状态
	Time        int64    `json:"time"`        //时间
	PhoneNumber int64    `json:"phoneNumber"` //手机号
	Content     string   `json:"content"`     //内容
	MediaIdList []string `json:"mediaIdList"` //图片 cdn 列表
	BlameResult int64    `json:"blameResult"` // 发货状态
}

// 投诉详细内容
type ComplaintOrderInfo struct {
	HeadImgUrl       string               `json:"headImgUrl"` //头像
	OrderId          string               `json:"orderId"`    //微信支付订单号
	OpenId           string               `json:"openId"`
	CreateTime       int64                `json:"createTime"`
	PhoneNumber      int64                `json:"phoneNumber"`
	Type             int64                `json:"type"`   //投诉问题分类
	Status           int64                `json:"status"` //订单状态，枚举值
	CustomerMaterial CustomerMaterialInfo `json:"customerMaterial"`
	ComplaintOrderId int64                `json:"complaintOrderId"`
	ProductName      string               `json:"productName"`  //商品名称
	PayTime          int64                `json:"payTime"`      //支付时间
	TotalCostStr     string               `json:"totalCostStr"` //金额
	ExpireTime       int64                `json:"expireTime"`   //过期时间
	OutTradeNo       string               `json:"outTradeNo"`   //过期时间
	NickName         string               `json:"nickName"`     //昵称

}
type CustomerMaterialInfo struct {
	Content     string   `json:"content"`
	MediaIdList []string `json:"mediaIdList"` // 投诉内容图片cdn(出于安全性考虑，目前该图片 url 有过期时间，如需查看图片，需每次重新调用该接口获取，防止图片过期导致无法查看)

}

type WxErrCodeReply struct {
	Errcode int64  `json:"errcode"` // 投诉内容
	Errmsg  string `json:"errmsg"`  // 投诉ID
}
type WxImage struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

type RequestQRCodeGetUnlimited struct {
	Scene     string  `json:"scene"`
	Page      string  `json:"page"`
	Width     int     `json:"width"`
	AutoColor bool    `json:"auto_color"`
	LineColor []int64 `json:"line_color"`
	IsHyaLine bool    `json:"is_hya_line"`
}

type MiniCode struct {
	Buffer      string `json:"buffer"`
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	ContentType string `json:"contentType"`
}
