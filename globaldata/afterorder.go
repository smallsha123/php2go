package globaldata



//售后状态
var SaleStatusMap = make(map[int64] string)
//售后原因
var RefundTypeMap  =   make(map[int64] string)
//售后类型
var SaleTypeMap = make(map[int64] string)
func init()  {
    SaleStatusMap[1]  = "待处理"
    //SaleStatusMap[2]  = "待卖家退货"
   // SaleStatusMap[3]  = "待商家收货"
    SaleStatusMap[6]  = "退款完成"
    SaleStatusMap[7]  = "售后关闭"
    SaleStatusMap[8]  = "退款中"
    SaleStatusMap[10]  = "处理完成"
    SaleStatusMap[12] = "退款失败"


    SaleTypeMap[1] = "仅退款"
    SaleTypeMap[2] = "退货退款"

    RefundTypeMap[1] = "退运费"
    RefundTypeMap[2] = "尺寸拍错/不喜欢/效果不好"
    RefundTypeMap[3] = "少发/漏发"
    RefundTypeMap[4] = "质量问题"
    RefundTypeMap[5] = "货物与描述不符"
    RefundTypeMap[6] = "未按约定时间发货"
    RefundTypeMap[7] = "发票问题"
    RefundTypeMap[8] = "卖家发错货"
    RefundTypeMap[9] = "假冒品牌"
    RefundTypeMap[10] = "退运费"
    RefundTypeMap[11] = "多拍/错拍/不想要"
    RefundTypeMap[12] = "空包裹"
    RefundTypeMap[13] = "快递/物流一直未送到"
    RefundTypeMap[14] = "快递/物流无跟踪记录"
    RefundTypeMap[15] = "货物破损已拒签"
    RefundTypeMap[16] = "其他"
    RefundTypeMap[17] = "用户取消订单"
    RefundTypeMap[18] = "商家取消订单"

}