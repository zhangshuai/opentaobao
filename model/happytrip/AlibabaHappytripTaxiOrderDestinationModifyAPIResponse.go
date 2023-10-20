package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiorderdestinationmodifyAPIResponse 修改目的地 API返回值
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
type AlibabahappytriptaxiorderdestinationmodifyAPIResponse struct {
	model.CommonResponse
	AlibabahappytriptaxiorderdestinationmodifyAPIResponseModel
}

// AlibabahappytriptaxiorderdestinationmodifyAPIResponseModel is 修改目的地 成功返回结果
type AlibabahappytriptaxiorderdestinationmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_destination_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误代码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}
