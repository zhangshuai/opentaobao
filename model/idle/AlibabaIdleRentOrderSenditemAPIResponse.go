package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentordersenditemAPIResponse 确认发货 API返回值
// alibaba.idle.rent.order.senditem
//
// 确认发货
type AlibabaidlerentordersenditemAPIResponse struct {
	model.CommonResponse
	AlibabaidlerentordersenditemAPIResponseModel
}

// AlibabaidlerentordersenditemAPIResponseModel is 确认发货 成功返回结果
type AlibabaidlerentordersenditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_senditem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaidlerentordersenditemTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
