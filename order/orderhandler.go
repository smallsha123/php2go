package orderx

const (
	SHOP_REMART    = 1  //商家修改备注
	USER_CANCEL    = -1 //订单取消
	ORDER_CREATE   = 2  //下单
	ORDER_PAY      = 3  //订单支付
	ORDER_SPLIT    = 4  // 订单发货
	ORDER_COMPLETE = 5  // 订单完成
	ORDER_EDIT     = 6  //订单编辑
)

func HandlerOrderNum(orderNum string) string {
	table := orderNum[:5]
	return "tk_order_" + table
}
