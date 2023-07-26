package wechatPay

// Params 基础参数
type Params struct {
	Name                          string  // 服务商全程
	MchID                         string  // 服务商商户号
	Appid                         string  // 服务商appid
	MchCertificateSerialNumber    string  // 商户证书序列号
	MchAPIv3Key                   string  // 商户APIv3密钥
	PrivateKey                    string  // 私钥oss地址
	PayNotify                     string  // 支付回调地址
	RefundNotify                  string  // 退款回调地址
	MerchantComplaintsNotify      string  // 微信商户支付投诉回调地址
	SplitPayRatio                 float64 // 支付时分账比例
	SplitPayDelayedSecond         int64   // 支付时分账延时秒数
	SplitConfirmDelayedSecond     int64   // 收货时分账延时秒数
	SplitQueryUpdateDelayedSecond int64   // 分账查询更新延时时间
}

// MerchantComplaintsUrlReply 商户投诉地址查询创建更新返回数据
type MerchantComplaintsUrlReply struct {
	Mchid string `json:"mchid"`
	URL   string `json:"url"`
}

// MerchantComplaintsUrlCreateReq 创建商户投诉地址请求数据
type MerchantComplaintsUrlCreateReq struct {
	URL string `json:"url"`
}

// MerchantComplaintsInfo 投诉详情
type MerchantComplaintsInfo struct {
	ComplaintDetail       string               `json:"complaint_detail"`
	ComplaintFullRefunded bool                 `json:"complaint_full_refunded"`
	ComplaintId           string               `json:"complaint_id"`
	ComplaintMediaList    []ComplaintMediaList `json:"complaint_media_list"`
	ComplaintOrderInfo    []ComplaintOrderInfo `json:"complaint_order_info"`
	ComplaintState        string               `json:"complaint_state"`
	ComplaintTime         string               `json:"complaint_time"`
	ComplaintedMchid      string               `json:"complainted_mchid"`
	IncomingUserResponse  bool                 `json:"incoming_user_response"`
	PayerOpenid           string               `json:"payer_openid"`
	PayerPhone            string               `json:"payer_phone"`
	ProblemDescription    string               `json:"problem_description"`
	ProblemType           string               `json:"problem_type"`
	ServiceOrderInfo      []ServiceOrderInfo   `json:"service_order_info"`
	UserComplaintTimes    int64                `json:"user_complaint_times"`
	UserTagList           []interface{}        `json:"user_tag_list"`
}

// ComplaintMediaList 投诉详情素材
type ComplaintMediaList struct {
	MediaType string   `json:"media_type"`
	MediaURL  []string `json:"media_url"`
}

// ComplaintOrderInfo 投诉详情订单信息
type ComplaintOrderInfo struct {
	Amount        int64  `json:"amount"`
	OutTradeNo    string `json:"out_trade_no"`
	TransactionId string `json:"transaction_id"`
}

// ServiceOrderInfo 投诉详情服务订单信息
type ServiceOrderInfo struct {
	OrderId    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

// MerchantComplaintsNotifyPlaintext 商户投诉回调明文
type MerchantComplaintsNotifyPlaintext struct {
	OutTradeNo           string `json:"out_trade_no"`
	ComplaintTime        string `json:"complaint_time"`
	Amount               int64  `json:"amount"`
	PayerPhone           string `json:"payer_phone"`
	ComplaintDetail      string `json:"complaint_detail"`
	ComplaintState       string `json:"complaint_state"`
	TransactionId        string `json:"transaction_id"`
	ComplaintHandleState string `json:"complaint_handle_state"`
	ActionType           string `json:"action_type"`
	ComplaintId          string `json:"complaint_id"`
}

// JsPayParam 服务商JsPay支付参数
type JsPayParam struct {
	SpAppid     string
	SpMchid     string
	SubAppid    string
	SubMchid    string
	Description string
	OutTradeNo  string
	Attach      string
	NotifyUrl   string
	GoodsTag    string
	LimitPay    []string
	Total       int64
	SpOpenid    string
	SubOpenid   string
	GoodsDetail *JsapiGoodsDetail
}

// JsapiGoodsDetail 服务商JsPay支付商品参数
type JsapiGoodsDetail struct {
	GoodsName        string
	MerchantGoodsId  string
	Quantity         int64
	UnitPrice        int64
	WechatpayGoodsId string
	InvoiceId        string
}

// PartnerRefundDomesticParams 服务商退款参数
type PartnerRefundDomesticParams struct {
	SubMchid      string
	TransactionId string
	OutTradeNo    string
	OutRefundNo   string
	Reason        string
	Refund        int64
	Total         int64
}

// PartnerRefundDomesticNotify 服务商退款回调
type PartnerRefundDomesticNotify struct {
	SPMchid             string                            `json:"sp_mchid"`
	SubMchid            string                            `json:"sub_mchid"`
	TransactionID       string                            `json:"transaction_id"`
	OutTradeNo          string                            `json:"out_trade_no"`
	RefundID            string                            `json:"refund_id"`
	OutRefundNo         string                            `json:"out_refund_no"`
	RefundStatus        string                            `json:"refund_status"`
	SuccessTime         string                            `json:"success_time"`
	UserReceivedAccount string                            `json:"user_received_account"`
	Amount              PartnerRefundDomesticNotifyAmount `json:"amount"`
}

// PartnerRefundDomesticNotifyAmount 服务窗退款回调金额
type PartnerRefundDomesticNotifyAmount struct {
	Total       int64 `json:"total"`
	Refund      int64 `json:"refund"`
	PayerTotal  int64 `json:"payer_total"`
	PayerRefund int64 `json:"payer_refund"`
}

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type PrepayWithRequestPaymentResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"` // revive:disable-line:var-naming
	// 应用ID
	Appid *string `json:"appId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}

// ProfitSharingCreateOrderRequest 分账创建订单请求
type ProfitSharingCreateOrderRequest struct {
	// 微信分配的服务商appid
	Appid string `json:"appid"`
	// 服务商系统内部的分账单号，在服务商系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo string `json:"out_order_no"`
	// 分账接收方列表，可以设置出资商户作为分账接受方，最多可有50个分账接收方
	Receivers []ProfitSharingCreateOrderReceiver `json:"receivers,omitempty"`
	// 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要）
	SubAppid string `json:"sub_appid,omitempty"`
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty"`
	// 微信支付订单号
	TransactionId string `json:"transaction_id"`
	// 1、如果为true，该笔订单剩余未分账的金额会解冻回分账方商户； 2、如果为false，该笔订单剩余未分账的金额不会解冻回分账方商户，可以对该笔订单再次进行分账。
	UnfreezeUnsplit bool `json:"unfreeze_unsplit"`
}

// ProfitSharingCreateOrderReceiver 分账创建订单请求
type ProfitSharingCreateOrderReceiver struct {
	// 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid  3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Account string `json:"account"`
	// 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Amount int64 `json:"amount"`
	// 分账的原因描述，分账账单中需要体现
	Description string `json:"description"`
	// 可选项，在接收方类型为个人的时可选填，若有值，会检查与 name 是否实名匹配，不匹配会拒绝分账请求 1、分账接收方类型是PERSONAL_OPENID或PERSONAL_SUB_OPENID时，是个人姓名的密文（选传，传则校验） 此字段的加密的方式为：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号
	Name string `json:"name,omitempty" encryption:"EM_APIV3"`
	// 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到） 3、PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）
	Type string `json:"type"`
}
