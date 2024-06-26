package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOnlineTicketnoGetAPIResponse 线上小票号获取 API返回值
// alibaba.mj.oc.online.ticketno.get
//
// 线上小票号获取
type AlibabaMjOcOnlineTicketnoGetAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcOnlineTicketnoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcOnlineTicketnoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcOnlineTicketnoGetAPIResponseModel).Reset()
}

// AlibabaMjOcOnlineTicketnoGetAPIResponseModel is 线上小票号获取 成功返回结果
type AlibabaMjOcOnlineTicketnoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_online_ticketno_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	ErrorsMsg string `json:"errors_msg,omitempty" xml:"errors_msg,omitempty"`
	// 错误码
	ErrorsCode int64 `json:"errors_code,omitempty" xml:"errors_code,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcOnlineTicketnoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.ErrorsMsg = ""
	m.ErrorsCode = 0
	m.IsSuccess = false
}

var poolAlibabaMjOcOnlineTicketnoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcOnlineTicketnoGetAPIResponse)
	},
}

// GetAlibabaMjOcOnlineTicketnoGetAPIResponse 从 sync.Pool 获取 AlibabaMjOcOnlineTicketnoGetAPIResponse
func GetAlibabaMjOcOnlineTicketnoGetAPIResponse() *AlibabaMjOcOnlineTicketnoGetAPIResponse {
	return poolAlibabaMjOcOnlineTicketnoGetAPIResponse.Get().(*AlibabaMjOcOnlineTicketnoGetAPIResponse)
}

// ReleaseAlibabaMjOcOnlineTicketnoGetAPIResponse 将 AlibabaMjOcOnlineTicketnoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcOnlineTicketnoGetAPIResponse(v *AlibabaMjOcOnlineTicketnoGetAPIResponse) {
	v.Reset()
	poolAlibabaMjOcOnlineTicketnoGetAPIResponse.Put(v)
}
