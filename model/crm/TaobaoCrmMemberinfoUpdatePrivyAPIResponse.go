package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberinfoUpdatePrivyAPIResponse 编辑会员资料 API返回值
// taobao.crm.memberinfo.update.privy
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
type TaobaoCrmMemberinfoUpdatePrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMemberinfoUpdatePrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMemberinfoUpdatePrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMemberinfoUpdatePrivyAPIResponseModel).Reset()
}

// TaobaoCrmMemberinfoUpdatePrivyAPIResponseModel is 编辑会员资料 成功返回结果
type TaobaoCrmMemberinfoUpdatePrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_memberinfo_update_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 会员信息修改是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMemberinfoUpdatePrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmMemberinfoUpdatePrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberinfoUpdatePrivyAPIResponse)
	},
}

// GetTaobaoCrmMemberinfoUpdatePrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmMemberinfoUpdatePrivyAPIResponse
func GetTaobaoCrmMemberinfoUpdatePrivyAPIResponse() *TaobaoCrmMemberinfoUpdatePrivyAPIResponse {
	return poolTaobaoCrmMemberinfoUpdatePrivyAPIResponse.Get().(*TaobaoCrmMemberinfoUpdatePrivyAPIResponse)
}

// ReleaseTaobaoCrmMemberinfoUpdatePrivyAPIResponse 将 TaobaoCrmMemberinfoUpdatePrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMemberinfoUpdatePrivyAPIResponse(v *TaobaoCrmMemberinfoUpdatePrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmMemberinfoUpdatePrivyAPIResponse.Put(v)
}
