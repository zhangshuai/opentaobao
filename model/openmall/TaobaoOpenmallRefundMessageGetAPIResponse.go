package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundMessageGetAPIResponse openmall获取退款单留言 API返回值
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
type TaobaoOpenmallRefundMessageGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundMessageGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundMessageGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundMessageGetAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundMessageGetAPIResponseModel is openmall获取退款单留言 成功返回结果
type TaobaoOpenmallRefundMessageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_message_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 留言列表
	ResultsList []RefundMessage `json:"results_list,omitempty" xml:"results_list>refund_message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundMessageGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultsList = m.ResultsList[:0]
}

var poolTaobaoOpenmallRefundMessageGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundMessageGetAPIResponse)
	},
}

// GetTaobaoOpenmallRefundMessageGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundMessageGetAPIResponse
func GetTaobaoOpenmallRefundMessageGetAPIResponse() *TaobaoOpenmallRefundMessageGetAPIResponse {
	return poolTaobaoOpenmallRefundMessageGetAPIResponse.Get().(*TaobaoOpenmallRefundMessageGetAPIResponse)
}

// ReleaseTaobaoOpenmallRefundMessageGetAPIResponse 将 TaobaoOpenmallRefundMessageGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundMessageGetAPIResponse(v *TaobaoOpenmallRefundMessageGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundMessageGetAPIResponse.Put(v)
}
