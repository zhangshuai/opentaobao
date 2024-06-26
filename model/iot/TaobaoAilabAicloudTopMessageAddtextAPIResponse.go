package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageAddtextAPIResponse 精灵代说 API返回值
// taobao.ailab.aicloud.top.message.addtext
//
// 精灵代说
type TaobaoAilabAicloudTopMessageAddtextAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessageAddtextAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageAddtextAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessageAddtextAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessageAddtextAPIResponseModel is 精灵代说 成功返回结果
type TaobaoAilabAicloudTopMessageAddtextAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_addtext_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageAddtextAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMessageAddtextAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageAddtextAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessageAddtextAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageAddtextAPIResponse
func GetTaobaoAilabAicloudTopMessageAddtextAPIResponse() *TaobaoAilabAicloudTopMessageAddtextAPIResponse {
	return poolTaobaoAilabAicloudTopMessageAddtextAPIResponse.Get().(*TaobaoAilabAicloudTopMessageAddtextAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessageAddtextAPIResponse 将 TaobaoAilabAicloudTopMessageAddtextAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageAddtextAPIResponse(v *TaobaoAilabAicloudTopMessageAddtextAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageAddtextAPIResponse.Put(v)
}
