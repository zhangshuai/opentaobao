package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityCreateAPIResponse 创建买赠活动 API返回值
// alibaba.retail.marketing.buygift.activity.create
//
// 同城供给买赠活动创建
type AlibabaRetailMarketingBuygiftActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivityCreateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivityCreateAPIResponseModel is 创建买赠活动 成功返回结果
type AlibabaRetailMarketingBuygiftActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivityCreateAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivityCreateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityCreateAPIResponse
func GetAlibabaRetailMarketingBuygiftActivityCreateAPIResponse() *AlibabaRetailMarketingBuygiftActivityCreateAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivityCreateAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivityCreateAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityCreateAPIResponse 将 AlibabaRetailMarketingBuygiftActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityCreateAPIResponse(v *AlibabaRetailMarketingBuygiftActivityCreateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityCreateAPIResponse.Put(v)
}
