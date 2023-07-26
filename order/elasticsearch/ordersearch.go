package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/smallsha123/php2go/errorx"
)

type (
	Demo struct {
		Shards       Shards        `json:"_shards"`
		Hits         Hits          `json:"hits"`
		TimedOut     bool          `json:"timed_out"`
		Took         int64         `json:"took"`
		Aggregations *Aggregations `json:"aggregations"`
	}
	Shards struct {
		Failed     int64 `json:"failed"`
		Skipped    int64 `json:"skipped"`
		Successful int64 `json:"successful"`
		Total      int64 `json:"total"`
	}
	Hits struct {
		Hits     []Hit   `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    Total   `json:"total"`
	}

	Total struct {
		Value    int64  `json:"value"`
		Relation string `json:"relation"`
	}
	Hit struct {
		ID      string  `json:"_id"`
		Index   string  `json:"_index"`
		Routing string  `json:"_routing"`
		Score   float64 `json:"_score"`
		Source  Source  `json:"_source"`
		Type    string  `json:"_type"`
	}
)

type Aggregations struct {
	OrderCount OrderCount `json:"order_count"`
}
type OrderCount struct {
	Buckets []Buckets `json:"buckets"`
}
type Buckets struct {
	Key      interface{} `json:"key"`
	DocCount int64       `json:"doc_count"`
}

type EsOrderList struct {
	Total        Total         `json:"total"`
	List         []EsOrderInfo `json:"hits"`
	Aggregations *Aggregations `json:"aggregations"`
}

type EsOrderInfo struct {
	Id     string `json:"_id"`
	Index  string `json:"_index"`
	Source Source `json:"_source"`
}
type Source struct {
	AdAccountId      int64  `json:"ad_account_id"`
	AdId             int64  `json:"ad_id"`
	AdminId          int64  `json:"admin_id"`
	AdminName        string `json:"admin_name"`
	AfterStatus      int64  `json:"after_status"`
	AppId            string `json:"app_id"`
	CallbackRatio    int64  `json:"callback_ratio"`
	CancelType       int64  `json:"cancel_type"`
	ChannelType      int64  `json:"channel_type"`
	CompanyId        int64  `json:"company_id"`
	CompleteTime     int64  `json:"complete_time"`
	CreatedAt        int64  `json:"created_at"`
	DeletedAt        int64  `json:"deleted_at"`
	DeliverTime      int64  `json:"deliver_time"`
	Discount         int64  `json:"discount"`
	FreightMoney     int64  `json:"freight_money"`
	Id               int64  `json:"id"`
	IsCallback       int64  `json:"is_callback"`
	IsOver           int64  `json:"is_over"`
	Money            int64  `json:"money"`
	OpenId           string `json:"open_id"`
	OrderId          int64  `json:"order_id"`
	OrderNum         string `json:"order_num"`
	OrderSource      int64  `json:"order_source"`
	PayMode          int64  `json:"pay_mode"`
	PayMoney         int64  `json:"pay_money"`
	PaySn            string `json:"pay_sn"`
	PayStateDesc     string `json:"pay_state_desc"`
	PayTime          int64  `json:"pay_time"`
	ReceiptTime      int64  `json:"receipt_time"`
	ResCallbackRatio int64  `json:"res_callback_ratio"`
	ShopId           int64  `json:"shop_id"`
	ShopRemark       string `json:"shop_remark"`
	SpMchId          string `json:"sp_mch_id"`
	SplitRun         int64  `json:"split_run"`
	SplitShopRatio   string `json:"split_shop_ratio"`
	SplitState       string `json:"split_state"`
	Status           int64  `json:"status"`
	SubMchId         string `json:"sub_mch_id"`
	TotalMoney       int64  `json:"total_money"`
	TransactionId    string `json:"transaction_id"`
	UpdatedAt        int64  `json:"updated_at"`
	UserId           int64  `json:"user_id"`
	UserRemark       string `json:"user_remark"`
	UtmSource        string `json:"utm_source"`
	Cid              string `json:"c_id"`
}

// GetOrderList query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"match": map[string]interface{}{
//				"_id": req.OrderId,
//			},
//		},
//	}
//	orderx.GetOrderList(query)

func GetOrderList(query map[string]interface{}) (*EsOrderList, error) {
	es, err := GetEsClient()
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error encoding query:"+err.Error())
	}
	var (
		r Demo
	)
	EsBodyLog(query)

	var buf bytes.Buffer

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error encoding query:"+err.Error())
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("tk_order"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error getting response:"+err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, " parsing the response body:"+err.Error())
		} else {
			// Print the response status and error information.
			return &EsOrderList{}, nil
			// return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, " parsing the response body:"+err.Error())
		}
	}

	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		// fmt.Println(err)
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, " 未查询到数据")
	}
	if err != nil {
		return nil, err
	}
	jsonRes, _ := json.Marshal(r.Hits)
	var orderList EsOrderList
	_ = json.Unmarshal(jsonRes, &orderList)
	if r.Aggregations != nil {
		aggrJson, _ := json.Marshal(r.Aggregations)
		_ = json.Unmarshal(aggrJson, &orderList.Aggregations)
	}
	return &orderList, nil
}
