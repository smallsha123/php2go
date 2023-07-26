package globalkey

/**
redis key except "model cache key"  in here,
but "model cache key" in model
*/

// CacheUserTokenKey /** 用户登陆的token
const CacheUserTokenKey = "user_token:%d"
const RedisLock = "redis_lock:%s"
const RedisShopLock = "redis_lock_shop:%d"
const RedisBossSmsCode = "redis_boss_sms_code_%s_%s"
const RedisBossSmsCodeLong = "redis_boss_sms_code_long_%s_%s"
const RedisBossSmsCodeOneHour = "redis_boss_sms_code_one_hour_%s_%s"
const RedisBossSmsCodeOneDaY = "redis_boss_sms_code_one_day_%s_%s"
const RedisOrderExpress = "order:express:%s:%s"
const RedisOrderBindInfo = "go:order:bind_info:%s"
const RedisProductShow = "go:product:show:%s"
const RedisShopProductShow = "go:product:shop:%s"

// RedisTopTypeProducts 首页数据缓存
const RedisTopTypeProducts = "go:top:%s:products:page:%s"

// RedisProductCategory 商品分类数据缓存
const RedisProductCategory = "go:product:%s:category"

// RedisOrderComplaintNum 投诉店铺数量缓存
const RedisOrderComplaintNum = "order:complaint:shop:%s"

// RedisHandleOrderComplaintNum 24小时处理投诉数量缓存
const RedisHandleOrderComplaintNum = "order:complaint:handle:shop:%d"

// RedisShopOrderNum 商铺订单数量
const RedisShopOrderNum = "shop:%d:order_num"

// RedisShopMerchantComplaintNum 商铺订单支付投诉数量
const RedisShopMerchantComplaintNum = "shop:%d:mch_complaint_num"

// RedisShopHandleMerchantComplaintNum 24小时处理商铺订单支付投诉数量
const RedisShopHandleMerchantComplaintNum = "shop:%d:handle_mch_complaint_num"

// RedisOrderNumGenerate 分布式订单号生成
const RedisOrderNumGenerate = "order:generate:machine_ids"
