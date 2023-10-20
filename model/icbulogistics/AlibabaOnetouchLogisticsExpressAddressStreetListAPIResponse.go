package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponse 四级地址库-街道 API返回值
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
type AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponse struct {
	model.CommonResponse
	AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponseModel
}

// AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponseModel is 四级地址库-街道 成功返回结果
type AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_street_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaonetouchlogisticsexpressaddressstreetlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
