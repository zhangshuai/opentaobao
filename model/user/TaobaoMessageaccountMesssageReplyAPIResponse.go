package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageReplyAPIResponse 消息号下行回复接口 API返回值
// taobao.messageaccount.messsage.reply
//
// 外部 isv 调用该进口来进行消息号消息的回复
type TaobaoMessageaccountMesssageReplyAPIResponse struct {
	model.CommonResponse
	TaobaoMessageaccountMesssageReplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageReplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMessageaccountMesssageReplyAPIResponseModel).Reset()
}

// TaobaoMessageaccountMesssageReplyAPIResponseModel is 消息号下行回复接口 成功返回结果
type TaobaoMessageaccountMesssageReplyAPIResponseModel struct {
	XMLName xml.Name `xml:"messageaccount_messsage_reply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoMessageaccountMesssageReplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMessageaccountMesssageReplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMessageaccountMesssageReplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMessageaccountMesssageReplyAPIResponse)
	},
}

// GetTaobaoMessageaccountMesssageReplyAPIResponse 从 sync.Pool 获取 TaobaoMessageaccountMesssageReplyAPIResponse
func GetTaobaoMessageaccountMesssageReplyAPIResponse() *TaobaoMessageaccountMesssageReplyAPIResponse {
	return poolTaobaoMessageaccountMesssageReplyAPIResponse.Get().(*TaobaoMessageaccountMesssageReplyAPIResponse)
}

// ReleaseTaobaoMessageaccountMesssageReplyAPIResponse 将 TaobaoMessageaccountMesssageReplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMessageaccountMesssageReplyAPIResponse(v *TaobaoMessageaccountMesssageReplyAPIResponse) {
	v.Reset()
	poolTaobaoMessageaccountMesssageReplyAPIResponse.Put(v)
}
