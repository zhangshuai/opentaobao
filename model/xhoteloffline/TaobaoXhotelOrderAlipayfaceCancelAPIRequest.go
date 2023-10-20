package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderalipayfacecancelAPIRequest 线下信用住订单取消 API请求
// taobao.xhotel.order.alipayface.cancel
//
// 提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
type TaobaoxhotelorderalipayfacecancelAPIRequest struct {
	model.Params
	// 原因描述
	_reasonText string
	// 外部订单号
	_outId string
	// 预留后续用
	_notifyUrl string
	// 请求流水号
	_outUuid string
	// 淘宝订单ID，必选
	_tid int64
}

// NewTaobaoxhotelorderalipayfacecancelRequest 初始化TaobaoxhotelorderalipayfacecancelAPIRequest对象
func NewTaobaoxhotelorderalipayfacecancelRequest() *TaobaoxhotelorderalipayfacecancelAPIRequest {
	return &TaobaoxhotelorderalipayfacecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReasonText is ReasonText Setter
// 原因描述
func (r *TaobaoxhotelorderalipayfacecancelAPIRequest) SetReasonText(_reasonText string) error {
	r._reasonText = _reasonText
	r.Set("reason_text", _reasonText)
	return nil
}

// GetReasonText ReasonText Getter
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetReasonText() string {
	return r._reasonText
}

// SetOutId is OutId Setter
// 外部订单号
func (r *TaobaoxhotelorderalipayfacecancelAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetOutId() string {
	return r._outId
}

// SetNotifyUrl is NotifyUrl Setter
// 预留后续用
func (r *TaobaoxhotelorderalipayfacecancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// GetNotifyUrl NotifyUrl Getter
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// SetOutUuid is OutUuid Setter
// 请求流水号
func (r *TaobaoxhotelorderalipayfacecancelAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetTid is Tid Setter
// 淘宝订单ID，必选
func (r *TaobaoxhotelorderalipayfacecancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelorderalipayfacecancelAPIRequest) GetTid() int64 {
	return r._tid
}
