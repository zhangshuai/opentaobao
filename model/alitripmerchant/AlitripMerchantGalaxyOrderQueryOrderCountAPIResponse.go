package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse 查询各种状态订单的总数 API返回值
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
type AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel is 查询各种状态订单的总数 成功返回结果
type AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_query_order_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderQueryOrderCountResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse
func GetAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse() *AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse {
	return poolAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse.Get().(*AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse 将 AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse(v *AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderQueryOrderCountAPIResponse.Put(v)
}
