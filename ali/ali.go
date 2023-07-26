package ali

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/smallsha123/php2go/globalkey"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type ExpressReply struct {
	Data    *ExpressInfo `json:"data"`
	Msg     string       `json:"msg"`
	Success bool         `json:"success"`
	Code    int          `json:"code"`
	TaskNo  string       `json:"taskNo"`
}
type ExpressInfo struct {
	ExpressCode           string           `json:"expressCode"`
	ExpressCompanyName    string           `json:"expressCompanyName"`
	Number                string           `json:"number"`
	LogisticsStatus       string           `json:"logisticsStatus"`
	LogisticsStatusDesc   string           `json:"logisticsStatusDesc"`
	TheLastMessage        string           `json:"theLastMessage"`
	TheLastTime           string           `json:"theLastTime"`
	TakeTime              string           `json:"takeTime"`
	LogisticsTraceDetails []*ExpressDetail `json:"logisticsTraceDetails"`
}

type ExpressDetail struct {
	AreaCode            string `json:"areaCode"`
	AreaName            string `json:"areaName"`
	SubLogisticsStatus  string `json:"subLogisticsStatus"`
	Time                int64  `json:"time"`
	LogisticsStatus     string `json:"logisticsStatus"`
	LogisticsStatusDesc string `json:"logisticsStatusDesc"`
	Desc                string `json:"desc"`
}

type CommonConf struct {
	AccessKeyId     string
	AccessKeySecret string
	RegionId        string
	Endpoint        string
}

type BaseSvc struct {
	*CommonConf
	config *openapi.Config
}

func NewBaseSvc(conf *CommonConf) *BaseSvc {
	config := new(openapi.Config)
	config.SetAccessKeyId(conf.AccessKeyId).
		SetAccessKeySecret(conf.AccessKeySecret).
		SetRegionId(conf.RegionId).SetEndpoint(conf.Endpoint)

	return &BaseSvc{
		CommonConf: conf,
		config:     config,
	}
}

// func commonConfig() *openapi.Config {
// 	config := new(openapi.Config)
// 	config.SetAccessKeyId("LTAI5t7Td5XxyNoq2zTLJXsD").
// 		SetAccessKeySecret("jbAweD3TA1m5vFx32peISLi7UFfN4S").
// 		SetRegionId("cn-hangzhou").SetEndpoint("afs.aliyuncs.com")
//
// 	return config
// }

func (s *BaseSvc) SendSms(phone, verifyCode string) error {
	client, err := sms.NewClient(s.config)
	if err != nil {
		return err
	}

	signName := "拓客引擎"
	templateCode := ""
	sendResp, err := client.SendSms(&sms.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &verifyCode,
	})
	if err != nil {
		logx.Errorf("[SendSms]验证码发送失败. phone:%s, err:%s", phone, err.Error())
		return err
	}

	b, _ := json.Marshal(sendResp)
	logx.Infof("[SendSms]验证码发送结果. phone:%s, resp:%s", phone, string(b))

	return nil
}

func QueryExpress(redis *redis.Redis, deliverySn string, mobile string) (infos *ExpressInfo, err error) {

	cacheExpressKey := fmt.Sprintf(globalkey.RedisOrderExpress, deliverySn, mobile)
	cacheExpressInfo, _ := redis.Get(cacheExpressKey)
	if cacheExpressInfo != "" {
		json.Unmarshal([]byte(cacheExpressInfo), &infos)
		return infos, nil
	}

	payload := strings.NewReader("number=" + deliverySn + "&mobile=" + mobile)

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://jmexpresv2.market.alicloudapi.com/express/query-v2", payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "APPCODE ce710b64756c45ed9ba816f53bf7947a")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := &ExpressReply{}
	json.Unmarshal(body, response)

	defExpressInfo := ExpressInfo{
		ExpressCode:         "",
		ExpressCompanyName:  "",
		LogisticsStatusDesc: "已揽件",
		TheLastMessage:      "快递已揽件！您的快递已经在路上",
		Number:              deliverySn,
	}
	defExpressInfo.LogisticsTraceDetails = append(defExpressInfo.LogisticsTraceDetails, &ExpressDetail{
		Desc:                "您的宝贝已经在路上了~",
		LogisticsStatusDesc: "已揽件",
	})
	if response.Code != 200 || reflect.DeepEqual(response.Data, ExpressReply{}.Data) {
		redisInfo, _ := json.Marshal(defExpressInfo)
		redis.Setex(cacheExpressKey, string(redisInfo), 30*60)
		return &defExpressInfo, nil
	}

	StatusMap := map[string]string{
		"WAIT_ACCEPT": "待揽收",
		"ACCEPT":      "已揽收",
		"TRANSPORT":   "运输中",
		"DELIVERING":  "派件中",
		"AGENT_SIGN":  "已代签收",
		"SIGN":        "已签收",
		"FAILED":      "包裹异常",
	}

	for _, detail := range response.Data.LogisticsTraceDetails {
		detail.LogisticsStatusDesc = StatusMap[detail.LogisticsStatus]
	}
	redisInfo, _ := json.Marshal(response.Data)
	redis.Setex(cacheExpressKey, string(redisInfo), 30*60)

	return response.Data, nil
}
