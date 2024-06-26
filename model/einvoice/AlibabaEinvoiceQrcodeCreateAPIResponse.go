package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceQrcodeCreateAPIResponse 扫码开票二维码生成 API返回值
// alibaba.einvoice.qrcode.create
//
// 扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
type AlibabaEinvoiceQrcodeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceQrcodeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceQrcodeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceQrcodeCreateAPIResponseModel).Reset()
}

// AlibabaEinvoiceQrcodeCreateAPIResponseModel is 扫码开票二维码生成 成功返回结果
type AlibabaEinvoiceQrcodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_qrcode_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaEinvoiceQrcodeCreateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceQrcodeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceQrcodeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceQrcodeCreateAPIResponse)
	},
}

// GetAlibabaEinvoiceQrcodeCreateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceQrcodeCreateAPIResponse
func GetAlibabaEinvoiceQrcodeCreateAPIResponse() *AlibabaEinvoiceQrcodeCreateAPIResponse {
	return poolAlibabaEinvoiceQrcodeCreateAPIResponse.Get().(*AlibabaEinvoiceQrcodeCreateAPIResponse)
}

// ReleaseAlibabaEinvoiceQrcodeCreateAPIResponse 将 AlibabaEinvoiceQrcodeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceQrcodeCreateAPIResponse(v *AlibabaEinvoiceQrcodeCreateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceQrcodeCreateAPIResponse.Put(v)
}
