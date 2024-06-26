package aliexpress

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialInsDirectresultUpdateAPIResponse ISV更新INS私信发送的结果 API返回值
// aliexpress.social.ins.directresult.update
//
// ISV更新INS私信发送的结果
type AliexpressSocialInsDirectresultUpdateAPIResponse struct {
	model.CommonResponse
	AliexpressSocialInsDirectresultUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialInsDirectresultUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialInsDirectresultUpdateAPIResponseModel).Reset()
}

// AliexpressSocialInsDirectresultUpdateAPIResponseModel is ISV更新INS私信发送的结果 成功返回结果
type AliexpressSocialInsDirectresultUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_ins_directresult_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorCodee string `json:"error_codee,omitempty" xml:"error_codee,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 更新是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 此次调用是否成功
	Successs bool `json:"successs,omitempty" xml:"successs,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialInsDirectresultUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCodee = ""
	m.ErrorMsg = ""
	m.Result = false
	m.Successs = false
}

var poolAliexpressSocialInsDirectresultUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialInsDirectresultUpdateAPIResponse)
	},
}

// GetAliexpressSocialInsDirectresultUpdateAPIResponse 从 sync.Pool 获取 AliexpressSocialInsDirectresultUpdateAPIResponse
func GetAliexpressSocialInsDirectresultUpdateAPIResponse() *AliexpressSocialInsDirectresultUpdateAPIResponse {
	return poolAliexpressSocialInsDirectresultUpdateAPIResponse.Get().(*AliexpressSocialInsDirectresultUpdateAPIResponse)
}

// ReleaseAliexpressSocialInsDirectresultUpdateAPIResponse 将 AliexpressSocialInsDirectresultUpdateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialInsDirectresultUpdateAPIResponse(v *AliexpressSocialInsDirectresultUpdateAPIResponse) {
	v.Reset()
	poolAliexpressSocialInsDirectresultUpdateAPIResponse.Put(v)
}
