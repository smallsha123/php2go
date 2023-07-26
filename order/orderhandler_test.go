package orderx

type AfterSaleListReq struct {
	ShopId          int64  `json:"shop_id"`
	RefundSn        string `json:"refund_sn"`
	OrderSn         string `json:"order_sn"`
	ProductName     string `json:"product_name"`
	SkuName         string `json:"sku_name"`
	AfterSaleStatus int64  `json:"after_sale_status"`
	AfterSaleType   int64  `json:"after_sale_type"`
	ReceiveName     string `json:"receive_name"`
	UserName        string `json:"user_name"`
	MsgId           int64  `json:"msg_id"`
	CreatedAt       string `json:"created_at"`
	AdminIds        string `json:"admin_ids"`
}
