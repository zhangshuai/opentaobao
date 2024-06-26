package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuBarcodeQueryAPIResponse 商品条码查询接口 API返回值
// alibaba.wdk.sku.barcode.query
//
// 查询商品编码，支持一品多码
type AlibabaWdkSkuBarcodeQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuBarcodeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuBarcodeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuBarcodeQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuBarcodeQueryAPIResponseModel is 商品条码查询接口 成功返回结果
type AlibabaWdkSkuBarcodeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_barcode_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuBarcodeQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuBarcodeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuBarcodeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuBarcodeQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuBarcodeQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuBarcodeQueryAPIResponse
func GetAlibabaWdkSkuBarcodeQueryAPIResponse() *AlibabaWdkSkuBarcodeQueryAPIResponse {
	return poolAlibabaWdkSkuBarcodeQueryAPIResponse.Get().(*AlibabaWdkSkuBarcodeQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuBarcodeQueryAPIResponse 将 AlibabaWdkSkuBarcodeQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuBarcodeQueryAPIResponse(v *AlibabaWdkSkuBarcodeQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuBarcodeQueryAPIResponse.Put(v)
}
