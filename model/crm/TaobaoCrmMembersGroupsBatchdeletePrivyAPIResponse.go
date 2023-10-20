package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse 批量删除分组（隐私号版） API返回值
// taobao.crm.members.groups.batchdelete.privy
//
// 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
type TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponseModel).Reset()
}

// TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponseModel is 批量删除分组（隐私号版） 成功返回结果
type TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_members_groups_batchdelete_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse)
	},
}

// GetTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse
func GetTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse() *TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse {
	return poolTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse.Get().(*TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse)
}

// ReleaseTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse 将 TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse(v *TaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmMembersGroupsBatchdeletePrivyAPIResponse.Put(v)
}
