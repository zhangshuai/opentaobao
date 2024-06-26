package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopUserSyncAPIResponse 外部人员同步 API返回值
// alitrip.btrip.corpop.user.sync
//
// 同步外部平台用户信息至商旅内部
type AlitripBtripCorpopUserSyncAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopUserSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopUserSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopUserSyncAPIResponseModel).Reset()
}

// AlitripBtripCorpopUserSyncAPIResponseModel is 外部人员同步 成功返回结果
type AlitripBtripCorpopUserSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_user_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回错误用户信息
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 成功标示
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopUserSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Module = ""
	m.ResultCode = 0
	m.SuccessFlag = false
}

var poolAlitripBtripCorpopUserSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopUserSyncAPIResponse)
	},
}

// GetAlitripBtripCorpopUserSyncAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopUserSyncAPIResponse
func GetAlitripBtripCorpopUserSyncAPIResponse() *AlitripBtripCorpopUserSyncAPIResponse {
	return poolAlitripBtripCorpopUserSyncAPIResponse.Get().(*AlitripBtripCorpopUserSyncAPIResponse)
}

// ReleaseAlitripBtripCorpopUserSyncAPIResponse 将 AlitripBtripCorpopUserSyncAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopUserSyncAPIResponse(v *AlitripBtripCorpopUserSyncAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopUserSyncAPIResponse.Put(v)
}
