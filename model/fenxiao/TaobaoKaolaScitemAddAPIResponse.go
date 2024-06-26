package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKaolaScitemAddAPIResponse 考拉货品新增接口 API返回值
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
type TaobaoKaolaScitemAddAPIResponse struct {
	model.CommonResponse
	TaobaoKaolaScitemAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKaolaScitemAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKaolaScitemAddAPIResponseModel).Reset()
}

// TaobaoKaolaScitemAddAPIResponseModel is 考拉货品新增接口 成功返回结果
type TaobaoKaolaScitemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"kaola_scitem_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常信息
	ErrorMessages []string `json:"error_messages,omitempty" xml:"error_messages>string,omitempty"`
	// 异常信息Code
	SysErrorCode string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`
	// 货品id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否系统异常
	IsSystemFailed bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKaolaScitemAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessages = m.ErrorMessages[:0]
	m.SysErrorCode = ""
	m.Data = 0
	m.IsSuccess = false
	m.IsSystemFailed = false
}

var poolTaobaoKaolaScitemAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKaolaScitemAddAPIResponse)
	},
}

// GetTaobaoKaolaScitemAddAPIResponse 从 sync.Pool 获取 TaobaoKaolaScitemAddAPIResponse
func GetTaobaoKaolaScitemAddAPIResponse() *TaobaoKaolaScitemAddAPIResponse {
	return poolTaobaoKaolaScitemAddAPIResponse.Get().(*TaobaoKaolaScitemAddAPIResponse)
}

// ReleaseTaobaoKaolaScitemAddAPIResponse 将 TaobaoKaolaScitemAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKaolaScitemAddAPIResponse(v *TaobaoKaolaScitemAddAPIResponse) {
	v.Reset()
	poolTaobaoKaolaScitemAddAPIResponse.Put(v)
}
