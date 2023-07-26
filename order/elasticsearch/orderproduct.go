package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/smallsha123/php2go/errorx"
)

type (
	Product struct {
		Shards       ProductShards `json:"_shards"`
		Hits         ProductHits   `json:"hits"`
		TimedOut     bool          `json:"timed_out"`
		Took         int64         `json:"took"`
		Aggregations *Aggregations `json:"aggregations"`
	}
	ProductShards struct {
		Failed     int64 `json:"failed"`
		Skipped    int64 `json:"skipped"`
		Successful int64 `json:"successful"`
		Total      int64 `json:"total"`
	}
	ProductHits struct {
		Hits     []ProductHit `json:"hits"`
		MaxScore float64      `json:"max_score"`
		Total    ProductTotal `json:"total"`
	}

	ProductTotal struct {
		Value    int64  `json:"value"`
		Relation string `json:"relation"`
	}
	ProductHit struct {
		ID      string        `json:"_id"`
		Index   string        `json:"_index"`
		Routing string        `json:"_routing"`
		Score   float64       `json:"_score"`
		Source  ProductSource `json:"_source"`
		Type    string        `json:"_type"`
	}
)
type EsOrderProductList struct {
	Total Total                `json:"total"`
	List  []EsOrderProductInfo `json:"hits"`
}
type EsOrderProductInfo struct {
	Id     string        `json:"_id"`
	Index  string        `json:"_index"`
	Source ProductSource `json:"_source"`
}
type ProductSource struct {
	Id          int64  `json:"id"`
	OrderId     int64  `json:"order_id"`
	ShopId      int64  `json:"shop_id"`
	ProductCode string `json:"product_code"` // 商品code
	SkuId       int64  `json:"sku_id"`       // 商品sukid
	SkuCode     string `json:"sku_code"`     // sku_code
	Price       int64  `json:"price"`        // 单价
	Num         int64  `json:"num"`          // 购买数量
	Image       string `json:"image"`        // 封面图
	Name        string `json:"name"`         // 商品名称
	SkuName     string `json:"sku_name"`     // 规格信息
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
	DeletedAt   int64  `json:"deleted_at"`   // 删除时间
}

func GetOrderProductList(query map[string]interface{}) (*EsOrderProductList, error) {
	es, err := GetEsClient()
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error encoding query:"+err.Error())
	}
	var (
		r Product
	)
	EsBodyLog(query)
	var buf bytes.Buffer

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "Error encoding query:"+err.Error())
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("tk_order_product"),
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
			return &EsOrderProductList{}, nil
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
	var orderProductList EsOrderProductList
	_ = json.Unmarshal(jsonRes, &orderProductList)
	// if r.Aggregations != nil {
	//	aggrJson, _ := json.Marshal(r.Aggregations)
	//	_ = json.Unmarshal(aggrJson, &orderProductList.Aggregations)
	// }
	return &orderProductList, nil
}
