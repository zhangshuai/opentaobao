package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIsvPageQueryAPIResponse isv离线会员数据分页查询 API返回值
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
type AlibabaMemberIsvPageQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMemberIsvPageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberIsvPageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberIsvPageQueryAPIResponseModel).Reset()
}

// AlibabaMemberIsvPageQueryAPIResponseModel is isv离线会员数据分页查询 成功返回结果
type AlibabaMemberIsvPageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_isv_page_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误E1002
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误明细
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回的分页结果
	Result *PageQueryIsvCustomerRes `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberIsvPageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Result = nil
}

var poolAlibabaMemberIsvPageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIsvPageQueryAPIResponse)
	},
}

// GetAlibabaMemberIsvPageQueryAPIResponse 从 sync.Pool 获取 AlibabaMemberIsvPageQueryAPIResponse
func GetAlibabaMemberIsvPageQueryAPIResponse() *AlibabaMemberIsvPageQueryAPIResponse {
	return poolAlibabaMemberIsvPageQueryAPIResponse.Get().(*AlibabaMemberIsvPageQueryAPIResponse)
}

// ReleaseAlibabaMemberIsvPageQueryAPIResponse 将 AlibabaMemberIsvPageQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberIsvPageQueryAPIResponse(v *AlibabaMemberIsvPageQueryAPIResponse) {
	v.Reset()
	poolAlibabaMemberIsvPageQueryAPIResponse.Put(v)
}
