package globaldata

import "os"

const (
	ExportDataBossQueue  = "tk:platform:api:common:export:data:boss"  // boss导出数据队列
	ExportDataAdminQueue = "tk:platform:api:common:export:data:admin" // admin导出数据队列
)

const (
	ExportDataOrderType         = "order"
	ExportDataAfterOrderType    = "refund"
	ExportDataDeliveryOrderType = "delivery"

	ExportDataShopType    = "shop"
	ExportDataCompanyType = "company"
)

const (
	ExportDataLocalSavePath = "./tmp/exportdata/"
)

var ExportDataBossType map[string]struct{}
var ExportDataAdminType map[string]struct{}

func init() {
	err := os.MkdirAll(ExportDataLocalSavePath, 0755)
	if err != nil {
		panic(err)
	}

	ExportDataBossType = make(map[string]struct{})
	ExportDataAdminType = make(map[string]struct{})

	// admin
	ExportDataAdminType[ExportDataOrderType] = struct{}{}
	ExportDataAdminType[ExportDataAfterOrderType] = struct{}{}
	ExportDataAdminType[ExportDataDeliveryOrderType] = struct{}{}

	// boss
	ExportDataBossType[ExportDataShopType] = struct{}{}
	ExportDataBossType[ExportDataCompanyType] = struct{}{}
}
