package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateAddAPIResponse 新增运费模板 API返回值
// taobao.delivery.template.add
//
// 增加运费模板的外部接口
type TaobaoDeliveryTemplateAddAPIResponse struct {
	model.CommonResponse
	TaobaoDeliveryTemplateAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeliveryTemplateAddAPIResponseModel).Reset()
}

// TaobaoDeliveryTemplateAddAPIResponseModel is 新增运费模板 成功返回结果
type TaobaoDeliveryTemplateAddAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 模板对象
	DeliveryTemplate *DeliveryTemplate `json:"delivery_template,omitempty" xml:"delivery_template,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryTemplate = nil
}

var poolTaobaoDeliveryTemplateAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeliveryTemplateAddAPIResponse)
	},
}

// GetTaobaoDeliveryTemplateAddAPIResponse 从 sync.Pool 获取 TaobaoDeliveryTemplateAddAPIResponse
func GetTaobaoDeliveryTemplateAddAPIResponse() *TaobaoDeliveryTemplateAddAPIResponse {
	return poolTaobaoDeliveryTemplateAddAPIResponse.Get().(*TaobaoDeliveryTemplateAddAPIResponse)
}

// ReleaseTaobaoDeliveryTemplateAddAPIResponse 将 TaobaoDeliveryTemplateAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeliveryTemplateAddAPIResponse(v *TaobaoDeliveryTemplateAddAPIResponse) {
	v.Reset()
	poolTaobaoDeliveryTemplateAddAPIResponse.Put(v)
}
