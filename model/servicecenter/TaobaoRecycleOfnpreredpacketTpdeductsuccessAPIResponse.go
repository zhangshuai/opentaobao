package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse 回收商同步前置补贴红包的代扣成功事件 API返回值
// taobao.recycle.ofnpreredpacket.tpdeductsuccess
//
// 回收商-&gt;天猫后端，同步前置补贴红包的代扣成功事件
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse struct {
	model.CommonResponse
	TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel).Reset()
}

// TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel is 回收商同步前置补贴红包的代扣成功事件 成功返回结果
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_ofnpreredpacket_tpdeductsuccess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作
	Data *OfnPreRedPacketActionDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse)
	},
}

// GetTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse 从 sync.Pool 获取 TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse
func GetTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse() *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse {
	return poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse.Get().(*TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse)
}

// ReleaseTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse 将 TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse(v *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse) {
	v.Reset()
	poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse.Put(v)
}
