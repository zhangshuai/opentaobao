package msgamp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageSendAPIRequest 消息发送 API请求
// taobao.message.send
//
// 消息发送接口
type TaobaoMessageSendAPIRequest struct {
	model.Params
	// 消息发送相关参数
	_sendMessageReq *SendMessageReq
}

// NewTaobaoMessageSendRequest 初始化TaobaoMessageSendAPIRequest对象
func NewTaobaoMessageSendRequest() *TaobaoMessageSendAPIRequest {
	return &TaobaoMessageSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMessageSendAPIRequest) Reset() {
	r._sendMessageReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMessageSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendMessageReq is SendMessageReq Setter
// 消息发送相关参数
func (r *TaobaoMessageSendAPIRequest) SetSendMessageReq(_sendMessageReq *SendMessageReq) error {
	r._sendMessageReq = _sendMessageReq
	r.Set("send_message_req", _sendMessageReq)
	return nil
}

// GetSendMessageReq SendMessageReq Getter
func (r TaobaoMessageSendAPIRequest) GetSendMessageReq() *SendMessageReq {
	return r._sendMessageReq
}

var poolTaobaoMessageSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMessageSendRequest()
	},
}

// GetTaobaoMessageSendRequest 从 sync.Pool 获取 TaobaoMessageSendAPIRequest
func GetTaobaoMessageSendAPIRequest() *TaobaoMessageSendAPIRequest {
	return poolTaobaoMessageSendAPIRequest.Get().(*TaobaoMessageSendAPIRequest)
}

// ReleaseTaobaoMessageSendAPIRequest 将 TaobaoMessageSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoMessageSendAPIRequest(v *TaobaoMessageSendAPIRequest) {
	v.Reset()
	poolTaobaoMessageSendAPIRequest.Put(v)
}
