package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockoutconfirmAPIResponse 出库单确认接口 API返回值
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
type TaobaoqimenstockoutconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoqimenstockoutconfirmAPIResponseModel
}

// TaobaoqimenstockoutconfirmAPIResponseModel is 出库单确认接口 成功返回结果
type TaobaoqimenstockoutconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockout_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenstockoutconfirmStruct `json:"response,omitempty" xml:"response,omitempty"`
}
