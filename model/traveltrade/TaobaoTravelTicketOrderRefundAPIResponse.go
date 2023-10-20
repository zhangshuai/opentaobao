package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotravelticketorderrefundAPIResponse 飞猪门票退票结果通知 API返回值
// taobao.travel.ticket.order.refund
//
// 门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。
type TaobaotravelticketorderrefundAPIResponse struct {
	model.CommonResponse
	TaobaotravelticketorderrefundAPIResponseModel
}

// TaobaotravelticketorderrefundAPIResponseModel is 飞猪门票退票结果通知 成功返回结果
type TaobaotravelticketorderrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"travel_ticket_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
