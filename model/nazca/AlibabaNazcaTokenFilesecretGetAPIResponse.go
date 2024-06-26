package nazca

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenFilesecretGetAPIResponse 获取文件秘钥 API返回值
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
type AlibabaNazcaTokenFilesecretGetAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaTokenFilesecretGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenFilesecretGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNazcaTokenFilesecretGetAPIResponseModel).Reset()
}

// AlibabaNazcaTokenFilesecretGetAPIResponseModel is 获取文件秘钥 成功返回结果
type AlibabaNazcaTokenFilesecretGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_filesecret_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件秘钥
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 错误信息
	SubErrorMessage string `json:"sub_error_message,omitempty" xml:"sub_error_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenFilesecretGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Error = ""
	m.Message = ""
	m.RetValue = ""
	m.SubErrorMessage = ""
	m.IsSuccess = false
}

var poolAlibabaNazcaTokenFilesecretGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNazcaTokenFilesecretGetAPIResponse)
	},
}

// GetAlibabaNazcaTokenFilesecretGetAPIResponse 从 sync.Pool 获取 AlibabaNazcaTokenFilesecretGetAPIResponse
func GetAlibabaNazcaTokenFilesecretGetAPIResponse() *AlibabaNazcaTokenFilesecretGetAPIResponse {
	return poolAlibabaNazcaTokenFilesecretGetAPIResponse.Get().(*AlibabaNazcaTokenFilesecretGetAPIResponse)
}

// ReleaseAlibabaNazcaTokenFilesecretGetAPIResponse 将 AlibabaNazcaTokenFilesecretGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNazcaTokenFilesecretGetAPIResponse(v *AlibabaNazcaTokenFilesecretGetAPIResponse) {
	v.Reset()
	poolAlibabaNazcaTokenFilesecretGetAPIResponse.Put(v)
}
