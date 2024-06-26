package aedata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateOrderListAPIResponse AE推广者订单批量获取接口 API返回值
// aliexpress.affiliate.order.list
//
// AE联盟推广者订单分页查询接口
type AliexpressAffiliateOrderListAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateOrderListAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateOrderListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateOrderListAPIResponseModel).Reset()
}

// AliexpressAffiliateOrderListAPIResponseModel is AE推广者订单批量获取接口 成功返回结果
type AliexpressAffiliateOrderListAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateOrderListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateOrderListAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateOrderListAPIResponse)
	},
}

// GetAliexpressAffiliateOrderListAPIResponse 从 sync.Pool 获取 AliexpressAffiliateOrderListAPIResponse
func GetAliexpressAffiliateOrderListAPIResponse() *AliexpressAffiliateOrderListAPIResponse {
	return poolAliexpressAffiliateOrderListAPIResponse.Get().(*AliexpressAffiliateOrderListAPIResponse)
}

// ReleaseAliexpressAffiliateOrderListAPIResponse 将 AliexpressAffiliateOrderListAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateOrderListAPIResponse(v *AliexpressAffiliateOrderListAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateOrderListAPIResponse.Put(v)
}
