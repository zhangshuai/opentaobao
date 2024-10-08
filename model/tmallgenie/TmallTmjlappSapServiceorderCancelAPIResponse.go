package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmjlappSapServiceorderCancelAPIResponse 取消售后服务单 API返回值
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
type TmallTmjlappSapServiceorderCancelAPIResponse struct {
	model.CommonResponse
	TmallTmjlappSapServiceorderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTmjlappSapServiceorderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTmjlappSapServiceorderCancelAPIResponseModel).Reset()
}

// TmallTmjlappSapServiceorderCancelAPIResponseModel is 取消售后服务单 成功返回结果
type TmallTmjlappSapServiceorderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmjlapp_sap_serviceorder_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取消服务单返回
	CancelResponse *Dtcancelresponse `json:"cancel_response,omitempty" xml:"cancel_response,omitempty"`
}

// Reset 清空结构体
func (m *TmallTmjlappSapServiceorderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CancelResponse = nil
}

var poolTmallTmjlappSapServiceorderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTmjlappSapServiceorderCancelAPIResponse)
	},
}

// GetTmallTmjlappSapServiceorderCancelAPIResponse 从 sync.Pool 获取 TmallTmjlappSapServiceorderCancelAPIResponse
func GetTmallTmjlappSapServiceorderCancelAPIResponse() *TmallTmjlappSapServiceorderCancelAPIResponse {
	return poolTmallTmjlappSapServiceorderCancelAPIResponse.Get().(*TmallTmjlappSapServiceorderCancelAPIResponse)
}

// ReleaseTmallTmjlappSapServiceorderCancelAPIResponse 将 TmallTmjlappSapServiceorderCancelAPIResponse 保存到 sync.Pool
func ReleaseTmallTmjlappSapServiceorderCancelAPIResponse(v *TmallTmjlappSapServiceorderCancelAPIResponse) {
	v.Reset()
	poolTmallTmjlappSapServiceorderCancelAPIResponse.Put(v)
}
