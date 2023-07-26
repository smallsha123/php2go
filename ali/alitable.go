package ali

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/smallsha123/php2go/globalkey"
	"strings"
	"time"
)

type ClientConfig struct {
	InstanceEndpoint string `json:"instance_endpoint"`
	InstanceName     string `json:"instance_name"`
	UserId           string `json:"user_id"`
	UserKey          string `json:"user_key"`
}

type BindInfo struct {
	Id           string `json:"id"`
	Sign         string `json:"sign"`
	AdminId      int64  `json:"admin_id"`
	AdId         int64  `json:"ad_id"`
	OpenId       string `json:"open_id"`
	UtmSource    *UtmSource
	CompanyId    int64  `json:"company_id"`
	UtmSourceStr string `json:"utm_source_str"`
	CreateAt     string `json:"create_at"`
}

type UtmSource struct {
	LogID   string `json:"log_id"`
	T       string `json:"t"`
	Clickid string `json:"clickid"`
	AdID    string `json:"ad_id"`
	Ua      string `json:"ua"`
}

type AppletJumpLog struct {
	Id         string `json:"id"`
	BindId     string `json:"bind_id"`
	AppId      string `json:"app_id"`
	Data       *UtmSource
	CreateTime int64 `json:"create_time"`
}

type ResInfo struct {
	Uid    string `json:"uid"`
	Aid    string `json:"aid"`
	OpenId string `json:"open_id"`
	Vid    string `json:"vid"`
}

func NewTableStore(params ClientConfig) *tablestore.TableStoreClient {
	client := tablestore.NewClient(params.InstanceEndpoint, params.InstanceName, params.UserId, params.UserKey)
	return client
}

func GetBindInfoByProduct(clientConfig ClientConfig, openId string, productId int64) (*BindInfo, error) {
	query := fmt.Sprintf("SELECT * FROM `tk_order_bind` where open_id = '%s' and product_id = %d and create_at > '%s' order by create_at desc LIMIT 1", openId, productId, time.Now().AddDate(0, 0, -7).Format(globalkey.DateTime))
	return GetBindInfoBySql(clientConfig, query)
}

func GetBindInfoBySign(clientConfig ClientConfig, openId string, sign string) (*BindInfo, error) {
	query := fmt.Sprintf("SELECT * FROM `tk_order_bind` where open_id = '%s' and sign = '%s' and create_at > '%s' order by create_at desc LIMIT 1", openId, sign, time.Now().AddDate(0, 0, -7).Format(globalkey.DateTime))
	return GetBindInfoBySql(clientConfig, query)
}

func GetBindInfoBySql(clientConfig ClientConfig, query string) (*BindInfo, error) {
	client := NewTableStore(clientConfig)
	response, err := client.SQLQuery(&tablestore.SQLQueryRequest{Query: query})
	if err != nil {
		return nil, err
	}
	resultSet := response.ResultSet
	if !resultSet.HasNext() {
		return nil, nil
	}

	bindInfo := &BindInfo{}
	for resultSet.HasNext() {
		row := resultSet.Next()
		bindInfo.Id, _ = row.GetStringByName("_id")
		bindInfo.Sign, _ = row.GetStringByName("sign")
		bindInfo.AdminId, _ = row.GetInt64ByName("admin_id")
		bindInfo.AdId, _ = row.GetInt64ByName("ad_id")
		bindInfo.OpenId, _ = row.GetStringByName("open_id")
		bindInfo.UtmSourceStr, _ = row.GetStringByName("utm_source")
		json.Unmarshal([]byte(bindInfo.UtmSourceStr), &bindInfo.UtmSource)
		bindInfo.CompanyId, _ = row.GetInt64ByName("company_id")
		bindInfo.CreateAt, _ = row.GetStringByName("create_at")
	}

	return bindInfo, nil
}

func GetAppletJumpInfo(clientConfig ClientConfig, bindId string) (*AppletJumpLog, error) {
	client := NewTableStore(clientConfig)
	query := fmt.Sprintf("SELECT * FROM `applet_jump_log` where bind_id = '%s' LIMIT 1", bindId)
	response, err := client.SQLQuery(&tablestore.SQLQueryRequest{Query: query})
	if err != nil {
		return nil, err
	}
	resultSet := response.ResultSet
	if !resultSet.HasNext() {
		return nil, nil
	}

	appletJumpLog := &AppletJumpLog{}
	for resultSet.HasNext() {
		row := resultSet.Next()
		appletJumpLog.Id, _ = row.GetStringByName("_id")
		appletJumpLog.BindId, _ = row.GetStringByName("bind_id")
		appletJumpLog.AppId, _ = row.GetStringByName("app_id")
		appletJumpLog.CreateTime, _ = row.GetInt64ByName("create_time")
		utmSource, _ := row.GetStringByName("data")
		json.Unmarshal([]byte(utmSource), &appletJumpLog.Data)
	}

	return appletJumpLog, nil
}

func SyncParams(source map[string]string) *ResInfo {

	resInfo := &ResInfo{}
	if gdtAdId, ok := source["gdt_ad_id"]; ok && gdtAdId != "" {
		resInfo.Aid = gdtAdId
	} else if wxAid, wxOk := source["wx_aid"]; wxOk && wxAid != "" {
		resInfo.Aid = wxAid
	} else if weiXinAdInfo, wxXinAdOk := source["weixinadinfo"]; wxXinAdOk && weiXinAdInfo != "" {
		dotIndex := strings.Index(weiXinAdInfo, ".")
		resInfo.Aid = weiXinAdInfo[:dotIndex]
	}

	if gdtVid, ok := source["gdt_vid"]; ok && gdtVid != "" {
		resInfo.Uid = gdtVid
	} else if qzGdt, qzGdtOk := source["qz_gdt"]; qzGdtOk && qzGdt != "" {
		resInfo.Uid = qzGdt
	} else if wxTraceId, wxTraceIdOk := source["wx_traceid"]; wxTraceIdOk && wxTraceId != "" {
		resInfo.Uid = wxTraceId
	}

	if openId, ok := source["open_id"]; ok && openId != "" {
		resInfo.OpenId = openId
	}

	return resInfo
}
