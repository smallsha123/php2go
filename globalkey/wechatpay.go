package globalkey

const (
	SplitTemp              = "split_temp"               // 支付后分账临时
	SplitPay               = "split_pay"                // 支付后请求分账
	SplitConfirm           = "split_confirm"            // 确认收货后请求分账
	SplitConfirmCommission = "split_confirm_commission" // 确认收货后请求分账分佣
	SplitUnfreeze          = "split_unfreeze"           // 解冻剩余资金
	SplitUnfreezeRefunds   = "split_unfreeze_refunds"   // 解冻剩余资金并且退款
	SplitQuery             = "split_query"              // 查询分账结果
	SplitQueryUpdate       = "split_query_update"       // 查询分账结果并更新数据库
	DomesticRefunds        = "domestic_refunds"         // 退款
)
