package icbuproduct

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductinventoryupdateAPIResponse icbu商品库存更新 API返回值
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
type AlibabaicbuproductinventoryupdateAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproductinventoryupdateAPIResponseModel
}

// AlibabaicbuproductinventoryupdateAPIResponseModel is icbu商品库存更新 成功返回结果
type AlibabaicbuproductinventoryupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
