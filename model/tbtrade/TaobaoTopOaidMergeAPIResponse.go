package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidMergeAPIResponse OAID订单合并 API返回值
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
type TaobaoTopOaidMergeAPIResponse struct {
	model.CommonResponse
	TaobaoTopOaidMergeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopOaidMergeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopOaidMergeAPIResponseModel).Reset()
}

// TaobaoTopOaidMergeAPIResponseModel is OAID订单合并 成功返回结果
type TaobaoTopOaidMergeAPIResponseModel struct {
	XMLName xml.Name `xml:"top_oaid_merge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 合单结果。注意：一个字符串中的多笔订单可以合单，不同字符串间的订单不可合单！！！&lt;br/&gt; 比如，[&#34;1111,2222&#34;]表示订单1111和订单2222可合单，[&#34;1111&#34;,&#34;2222&#34;]表示订单1111和订单2222不可以合单。
	TidList []string `json:"tid_list,omitempty" xml:"tid_list>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopOaidMergeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TidList = m.TidList[:0]
}

var poolTaobaoTopOaidMergeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopOaidMergeAPIResponse)
	},
}

// GetTaobaoTopOaidMergeAPIResponse 从 sync.Pool 获取 TaobaoTopOaidMergeAPIResponse
func GetTaobaoTopOaidMergeAPIResponse() *TaobaoTopOaidMergeAPIResponse {
	return poolTaobaoTopOaidMergeAPIResponse.Get().(*TaobaoTopOaidMergeAPIResponse)
}

// ReleaseTaobaoTopOaidMergeAPIResponse 将 TaobaoTopOaidMergeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopOaidMergeAPIResponse(v *TaobaoTopOaidMergeAPIResponse) {
	v.Reset()
	poolTaobaoTopOaidMergeAPIResponse.Put(v)
}
