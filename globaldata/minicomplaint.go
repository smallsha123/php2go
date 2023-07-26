package globaldata

//售后状态
var  ComplaintType = make(map[int64] string)
//售后原因
var ComplaintStatus  =   make(map[int64] string)
//小程序售后
var  MiniComplaintType = make(map[int64] string)
//投诉状态数据库搜索
var MiniStatus =make(map[int64] []int64)
func init()  {

    ComplaintType[1] = "发货问题"
    ComplaintType[2] = "商品问题"
    ComplaintType[3] = "客服问题"
    ComplaintType[4] = "其他问题"
    ComplaintType[611] = "未按约定时间发货"
    ComplaintType[612] = "商家拒绝发货"
    ComplaintType[613] = "少发/错发"
    ComplaintType[614] = "物流信息长时间不更新"
    ComplaintType[621] = "客服不回复"
    ComplaintType[622] = "客服辱骂/骚扰/恐吓"
    ComplaintType[631] = "赠品承诺未履行"
    ComplaintType[632] = "物流承诺未履行"
    ComplaintType[633] = "承诺未履行其他"
    ComplaintType[641] = "描述不符"
    ComplaintType[642] = "商品破损"
    ComplaintType[643] = "商品其他"

    //日志状态时间轴
    ComplaintStatus[1] = "待处理"
    ComplaintStatus[2] = "处理中"
    ComplaintStatus[3] = "处理完成"
   // ComplaintStatus[4] = "处理异常"
    ComplaintStatus[5] = "拒绝和解"
    MiniComplaintType[1] = "用户发起投诉"
    MiniComplaintType[2] = "用户补充留言"
    MiniComplaintType[3] = "商家补充留言"
    MiniComplaintType[7] = "用户补充凭证"
    MiniComplaintType[8] = "商家补充凭证"
    MiniComplaintType[11] = "用户申请平台客服协助"
    MiniComplaintType[12] = "用户撤销投诉"
    MiniComplaintType[13] = "平台客服处理中"
    MiniComplaintType[14] = "待用户补充凭证"
    MiniComplaintType[16] = "待商家补充凭证"
    MiniComplaintType[18] = "平台要求双方补充凭证"
    MiniComplaintType[26] = "平台核实处理凭证异常，投诉关闭，请商家自行联系用户解决问题，保障用户体验"
    MiniComplaintType[30] = "平台已核实此投诉非商家责任，投诉已完结"
    MiniComplaintType[31] = "平台已核实此投诉为商家责任，待上传处理凭证（blameResult为0）/平台已核实此投诉为商家责任，待用户退货中（blameResult为1）"
    MiniComplaintType[32] = "平台已核实此投诉为商家责任，待上传处理凭证（blameResult为0）/平台已核实此投诉为商家责任，待用户退货中（blameResult为1）"
    MiniComplaintType[33] = "平台已核实此投诉非商家责任，投诉已完结"
    MiniComplaintType[36] = "平台已核实处理凭证，投诉完结"
    MiniComplaintType[37] = "平台核实处理凭证异常，投诉关闭，请商家自行联系用户解决问题，保障用户体验"
    MiniComplaintType[101] = "商家超时未回应投诉"
    MiniComplaintType[104] = "用户认可处理结果，投诉已完结"
    MiniComplaintType[107] = "商家超时未提交投诉处理凭证，平台客服处理中"
    MiniComplaintType[108] = "用户超时未确认商家回应结果，投诉已完结"
    MiniComplaintType[109] = "商家已回应投诉"
    MiniComplaintType[110] = "商家提交投诉处理凭证"
    MiniComplaintType[111] = "用户补充凭证超时"
    MiniComplaintType[112] = "商家补充凭证超时"
    MiniComplaintType[113] = "双方补充凭证超时"


    //MiniComplaintType[101] = "平台客服处理中"
    //MiniComplaintType[102] = "用户取消申请"
    //MiniComplaintType[103] = "平台客服处理中"
    //MiniComplaintType[104] = "平台客服处理中"
    //MiniComplaintType[105] = "平台客服处理中"
    //MiniComplaintType[106] = "待商家补充凭证"
    //MiniComplaintType[107] = "平台客服处理中"
    //MiniComplaintType[108] = "待双方补充凭证"
    //MiniComplaintType[109] = "平台客服处理中"
    //MiniComplaintType[112] = "投诉已完结"
    //MiniComplaintType[115] = "投诉已完结"
    //MiniComplaintType[116] = "投诉已完结"
    //MiniComplaintType[201] = "待处理"
    //MiniComplaintType[202] = "商家超时未回应，待用户确认"
    //MiniComplaintType[203] = "已回应,待用户确认"
    //MiniComplaintType[204] = "已回应,待用户确认"
    //MiniComplaintType[205] = "投诉已完结"
    //MiniComplaintType[206] = "平台已判定为商责，待上传处理凭证"
    //MiniComplaintType[207] = "平台客服核实凭证中"
    //MiniComplaintType[208] = "超时未上传凭证"
    //MiniComplaintType[209] = "投诉已关闭"
    //MiniComplaintType[305] = "平台客服处理中"
    //MiniComplaintType[307] = "平台客服处理中"
    //MiniComplaintType[308] = "平台已判定为商责，待用户退货中"
    //MiniComplaintType[309] = "平台已判定为商责，待用户退货中"
    //MiniComplaintType[310] = "平台客服处理中"
    //MiniComplaintType[311] = "签收异常"

    MiniStatus[101] = []int64{101,103,105,107,109,305,307,310}
    MiniStatus[112] = []int64{112,115,116,205}
    MiniStatus[204] = []int64{203,204}
    MiniStatus[308] = []int64{308,309}
}