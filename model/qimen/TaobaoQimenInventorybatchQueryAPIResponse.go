package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventorybatchqueryAPIResponse 商品单仓批次库存查询接口 API返回值
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
type TaobaoqimeninventorybatchqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimeninventorybatchqueryAPIResponseModel
}

// TaobaoqimeninventorybatchqueryAPIResponseModel is 商品单仓批次库存查询接口 成功返回结果
type TaobaoqimeninventorybatchqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventorybatch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Response *TaobaoqimeninventorybatchqueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
