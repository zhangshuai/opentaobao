package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundGetAPIResponse 获取OpenMall退款单详情 API返回值
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
type TaobaoOpenmallRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundGetAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundGetAPIResponseModel is 获取OpenMall退款单详情 成功返回结果
type TaobaoOpenmallRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Refund *TopRefundVo `json:"refund,omitempty" xml:"refund,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refund = nil
}

var poolTaobaoOpenmallRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundGetAPIResponse)
	},
}

// GetTaobaoOpenmallRefundGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundGetAPIResponse
func GetTaobaoOpenmallRefundGetAPIResponse() *TaobaoOpenmallRefundGetAPIResponse {
	return poolTaobaoOpenmallRefundGetAPIResponse.Get().(*TaobaoOpenmallRefundGetAPIResponse)
}

// ReleaseTaobaoOpenmallRefundGetAPIResponse 将 TaobaoOpenmallRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundGetAPIResponse(v *TaobaoOpenmallRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundGetAPIResponse.Put(v)
}
