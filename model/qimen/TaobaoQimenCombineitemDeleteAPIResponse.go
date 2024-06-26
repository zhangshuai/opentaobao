package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemDeleteAPIResponse 组合货品删除接口 API返回值
// taobao.qimen.combineitem.delete
//
// 组合货品删除
type TaobaoQimenCombineitemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoQimenCombineitemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenCombineitemDeleteAPIResponseModel).Reset()
}

// TaobaoQimenCombineitemDeleteAPIResponseModel is 组合货品删除接口 成功返回结果
type TaobaoQimenCombineitemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_combineitem_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenCombineitemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemDeleteAPIResponse)
	},
}

// GetTaobaoQimenCombineitemDeleteAPIResponse 从 sync.Pool 获取 TaobaoQimenCombineitemDeleteAPIResponse
func GetTaobaoQimenCombineitemDeleteAPIResponse() *TaobaoQimenCombineitemDeleteAPIResponse {
	return poolTaobaoQimenCombineitemDeleteAPIResponse.Get().(*TaobaoQimenCombineitemDeleteAPIResponse)
}

// ReleaseTaobaoQimenCombineitemDeleteAPIResponse 将 TaobaoQimenCombineitemDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenCombineitemDeleteAPIResponse(v *TaobaoQimenCombineitemDeleteAPIResponse) {
	v.Reset()
	poolTaobaoQimenCombineitemDeleteAPIResponse.Put(v)
}
