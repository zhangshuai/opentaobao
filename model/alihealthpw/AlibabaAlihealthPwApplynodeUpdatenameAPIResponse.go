package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwApplynodeUpdatenameAPIResponse 回调变更患者姓名 API返回值
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
type AlibabaAlihealthPwApplynodeUpdatenameAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwApplynodeUpdatenameAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwApplynodeUpdatenameAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwApplynodeUpdatenameAPIResponseModel).Reset()
}

// AlibabaAlihealthPwApplynodeUpdatenameAPIResponseModel is 回调变更患者姓名 成功返回结果
type AlibabaAlihealthPwApplynodeUpdatenameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_applynode_updatename_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pap项目状态码
	PapCode string `json:"pap_code,omitempty" xml:"pap_code,omitempty"`
	// pap项目状态描述
	PapMessage string `json:"pap_message,omitempty" xml:"pap_message,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwApplynodeUpdatenameAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PapCode = ""
	m.PapMessage = ""
}

var poolAlibabaAlihealthPwApplynodeUpdatenameAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwApplynodeUpdatenameAPIResponse)
	},
}

// GetAlibabaAlihealthPwApplynodeUpdatenameAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwApplynodeUpdatenameAPIResponse
func GetAlibabaAlihealthPwApplynodeUpdatenameAPIResponse() *AlibabaAlihealthPwApplynodeUpdatenameAPIResponse {
	return poolAlibabaAlihealthPwApplynodeUpdatenameAPIResponse.Get().(*AlibabaAlihealthPwApplynodeUpdatenameAPIResponse)
}

// ReleaseAlibabaAlihealthPwApplynodeUpdatenameAPIResponse 将 AlibabaAlihealthPwApplynodeUpdatenameAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwApplynodeUpdatenameAPIResponse(v *AlibabaAlihealthPwApplynodeUpdatenameAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwApplynodeUpdatenameAPIResponse.Put(v)
}
