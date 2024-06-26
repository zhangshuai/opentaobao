package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelCrsorderSearchAPIResponse CRS接送机订单列表搜索 API返回值
// alitrip.travel.crsorder.search
//
// 提供给CRS商家搜索订单列表，仅返回订单号列表
type AlitripTravelCrsorderSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelCrsorderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelCrsorderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelCrsorderSearchAPIResponseModel).Reset()
}

// AlitripTravelCrsorderSearchAPIResponseModel is CRS接送机订单列表搜索 成功返回结果
type AlitripTravelCrsorderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_crsorder_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单id列表（string类型）
	OrderStringList []string `json:"order_string_list,omitempty" xml:"order_string_list>string,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelCrsorderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderStringList = m.OrderStringList[:0]
}

var poolAlitripTravelCrsorderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelCrsorderSearchAPIResponse)
	},
}

// GetAlitripTravelCrsorderSearchAPIResponse 从 sync.Pool 获取 AlitripTravelCrsorderSearchAPIResponse
func GetAlitripTravelCrsorderSearchAPIResponse() *AlitripTravelCrsorderSearchAPIResponse {
	return poolAlitripTravelCrsorderSearchAPIResponse.Get().(*AlitripTravelCrsorderSearchAPIResponse)
}

// ReleaseAlitripTravelCrsorderSearchAPIResponse 将 AlitripTravelCrsorderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelCrsorderSearchAPIResponse(v *AlitripTravelCrsorderSearchAPIResponse) {
	v.Reset()
	poolAlitripTravelCrsorderSearchAPIResponse.Put(v)
}
