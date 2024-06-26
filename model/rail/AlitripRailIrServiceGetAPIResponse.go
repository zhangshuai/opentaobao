package rail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrServiceGetAPIResponse 国际火车票仓位坐席查询 API返回值
// alitrip.rail.ir.service.get
//
// 国际火车票标准仓位坐席查询
type AlitripRailIrServiceGetAPIResponse struct {
	model.CommonResponse
	AlitripRailIrServiceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripRailIrServiceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripRailIrServiceGetAPIResponseModel).Reset()
}

// AlitripRailIrServiceGetAPIResponseModel is 国际火车票仓位坐席查询 成功返回结果
type AlitripRailIrServiceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_ir_service_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripRailIrServiceGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripRailIrServiceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripRailIrServiceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripRailIrServiceGetAPIResponse)
	},
}

// GetAlitripRailIrServiceGetAPIResponse 从 sync.Pool 获取 AlitripRailIrServiceGetAPIResponse
func GetAlitripRailIrServiceGetAPIResponse() *AlitripRailIrServiceGetAPIResponse {
	return poolAlitripRailIrServiceGetAPIResponse.Get().(*AlitripRailIrServiceGetAPIResponse)
}

// ReleaseAlitripRailIrServiceGetAPIResponse 将 AlitripRailIrServiceGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripRailIrServiceGetAPIResponse(v *AlitripRailIrServiceGetAPIResponse) {
	v.Reset()
	poolAlitripRailIrServiceGetAPIResponse.Put(v)
}
