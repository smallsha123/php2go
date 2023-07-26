package errorx

var message = make(map[uint32]string)

const (
	OK                         = 0    // 成功
	ERR_DEFAULT                = 1    // 错误
	ERR_EMPTY                  = 400  // 未查到数据
	ERR_UNAUTHORIZED           = 401  // 未授权
	ERR_PARAMS                 = 499  // 参数错误
	ERR_SHOP_EXIST             = 1000 // 店铺已存在
	ERR_SKU_EMPTY              = 1001 // sku信息不能为空
	ERR_PRODUCT_FOUND          = 1002 // 商品不存在
	ERR_NOT_RECERIVE_AD_COUPON = 1003 // 没有可领取的广告优惠券
	ERR_ORDER_PRODUCT_UN_SALE  = 1004 // 商品或Sku下架
	ERR_ORDER_MONEY            = 1005 // 订单金额异常
	ERR_ORDER_ADDRESS_DISABLE  = 1006 // 订单地址不可用

)

func init() {
	message[OK] = "操作成功"
	message[ERR_DEFAULT] = "系统未知错误"
	message[ERR_EMPTY] = "未查到数据"
	message[ERR_UNAUTHORIZED] = "unauthorized"
	message[ERR_PARAMS] = "参数错误"
	message[ERR_SHOP_EXIST] = "店铺名称不可重复"
	message[ERR_SKU_EMPTY] = "sku信息不能为空"
	message[ERR_PRODUCT_FOUND] = "商品不存在"
	message[ERR_NOT_RECERIVE_AD_COUPON] = "没有可领取的广告优惠券"
}
