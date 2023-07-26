package push

// 订单通知
const (
	OrderNew             = "Y9v2pRhtLxo9fEH09ftgXQdQd1Re4dSG" // 新订单通知
	OrderWaitDelivery    = "mrTkHKzsQABAC3P3oWmnyWmC2cvR2oSg" // 待发货
	OrderDeliveryTimeout = "DdmYmv2l1EEfvPaVQKNfAJ2Hp1bY3Mwq" // 待发货即将超时
)

// 投诉通知
const (
	ComplainWechatPay = "jZyC8i6kPkbEFc4iub2WfJte7JL73PtM" // 新投诉-微信支付投
	ComplainShop      = "WDE99PrpbvfigqWWNF93AD9YeCNW8PwF" // 新投诉-店铺投诉
)

// 售后服务
const (
	OfterSaleWaitHandle = "eAfXFkVEqjTIYKrqGRsn0i8tedrgi6Hy" // 待处理售后
	OfterSaleWaitTake   = "P2A8Z0suxzDPATmcqL73dX6gSKnBxpAu" // 待收货
)

// 账号通知

// 商品通知
const (
	ProductExamineFail        = "wo83h7I161JMokrJxGeL2rXWZmOo3oIH" // 商品审核不通过
	ProductViolationDownShelf = "xaR5OBMYvzLxkk7yccMgM2QrQRyna5tr" // 商品违规下架
)
