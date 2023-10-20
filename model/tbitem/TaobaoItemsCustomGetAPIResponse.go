package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemsCustomGetAPIResponse 根据外部ID取商品 API返回值
// taobao.items.custom.get
//
// 跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemsCustomGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemsCustomGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemsCustomGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemsCustomGetAPIResponseModel).Reset()
}

// TaobaoItemsCustomGetAPIResponseModel is 根据外部ID取商品 成功返回结果
type TaobaoItemsCustomGetAPIResponseModel struct {
	XMLName xml.Name `xml:"items_custom_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品列表，具体返回字段以fields决定
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemsCustomGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Items = m.Items[:0]
}

var poolTaobaoItemsCustomGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemsCustomGetAPIResponse)
	},
}

// GetTaobaoItemsCustomGetAPIResponse 从 sync.Pool 获取 TaobaoItemsCustomGetAPIResponse
func GetTaobaoItemsCustomGetAPIResponse() *TaobaoItemsCustomGetAPIResponse {
	return poolTaobaoItemsCustomGetAPIResponse.Get().(*TaobaoItemsCustomGetAPIResponse)
}

// ReleaseTaobaoItemsCustomGetAPIResponse 将 TaobaoItemsCustomGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemsCustomGetAPIResponse(v *TaobaoItemsCustomGetAPIResponse) {
	v.Reset()
	poolTaobaoItemsCustomGetAPIResponse.Put(v)
}
