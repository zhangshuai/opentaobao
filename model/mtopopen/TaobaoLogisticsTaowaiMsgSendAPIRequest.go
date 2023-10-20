package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticstaowaimsgsendAPIRequest 淘外包裹物流信息走淘宝发包裹状态通知接口 API请求
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
type TaobaologisticstaowaimsgsendAPIRequest struct {
	model.Params
	// 请求对象
	_taowaiMsgSendRequest *TaowaiMsgSendRequest
}

// NewTaobaologisticstaowaimsgsendRequest 初始化TaobaologisticstaowaimsgsendAPIRequest对象
func NewTaobaologisticstaowaimsgsendRequest() *TaobaologisticstaowaimsgsendAPIRequest {
	return &TaobaologisticstaowaimsgsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticstaowaimsgsendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.taowai.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticstaowaimsgsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticstaowaimsgsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaowaiMsgSendRequest is TaowaiMsgSendRequest Setter
// 请求对象
func (r *TaobaologisticstaowaimsgsendAPIRequest) SetTaowaiMsgSendRequest(_taowaiMsgSendRequest *TaowaiMsgSendRequest) error {
	r._taowaiMsgSendRequest = _taowaiMsgSendRequest
	r.Set("taowai_msg_send_request", _taowaiMsgSendRequest)
	return nil
}

// GetTaowaiMsgSendRequest TaowaiMsgSendRequest Getter
func (r TaobaologisticstaowaimsgsendAPIRequest) GetTaowaiMsgSendRequest() *TaowaiMsgSendRequest {
	return r._taowaiMsgSendRequest
}