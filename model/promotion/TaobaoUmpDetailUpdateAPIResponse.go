package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailUpdateAPIResponse 修改活动详情 API返回值
// taobao.ump.detail.update
//
// 更新活动详情
type TaobaoUmpDetailUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoUmpDetailUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpDetailUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpDetailUpdateAPIResponseModel).Reset()
}

// TaobaoUmpDetailUpdateAPIResponseModel is 修改活动详情 成功返回结果
type TaobaoUmpDetailUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpDetailUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoUmpDetailUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpDetailUpdateAPIResponse)
	},
}

// GetTaobaoUmpDetailUpdateAPIResponse 从 sync.Pool 获取 TaobaoUmpDetailUpdateAPIResponse
func GetTaobaoUmpDetailUpdateAPIResponse() *TaobaoUmpDetailUpdateAPIResponse {
	return poolTaobaoUmpDetailUpdateAPIResponse.Get().(*TaobaoUmpDetailUpdateAPIResponse)
}

// ReleaseTaobaoUmpDetailUpdateAPIResponse 将 TaobaoUmpDetailUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpDetailUpdateAPIResponse(v *TaobaoUmpDetailUpdateAPIResponse) {
	v.Reset()
	poolTaobaoUmpDetailUpdateAPIResponse.Put(v)
}
