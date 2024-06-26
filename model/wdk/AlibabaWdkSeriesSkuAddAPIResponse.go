package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSkuAddAPIResponse 系列品商品变更-添加商品 API返回值
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
type AlibabaWdkSeriesSkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesSkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesSkuAddAPIResponseModel).Reset()
}

// AlibabaWdkSeriesSkuAddAPIResponseModel is 系列品商品变更-添加商品 成功返回结果
type AlibabaWdkSeriesSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesSkuAddApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesSkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesSkuAddAPIResponse)
	},
}

// GetAlibabaWdkSeriesSkuAddAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesSkuAddAPIResponse
func GetAlibabaWdkSeriesSkuAddAPIResponse() *AlibabaWdkSeriesSkuAddAPIResponse {
	return poolAlibabaWdkSeriesSkuAddAPIResponse.Get().(*AlibabaWdkSeriesSkuAddAPIResponse)
}

// ReleaseAlibabaWdkSeriesSkuAddAPIResponse 将 AlibabaWdkSeriesSkuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesSkuAddAPIResponse(v *AlibabaWdkSeriesSkuAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesSkuAddAPIResponse.Put(v)
}
