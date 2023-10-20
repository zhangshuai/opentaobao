package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripshipordernotifyAPIRequest 订单信息回填(出票回调) API请求
// alitrip.ship.order.notify
//
// 此接口为接入商调用飞猪旅行接口回填票号、密码(验证码)等订单信息。接口根据alitripOrderId幂等。若第一次调用失败，后续调用仍然可以回填票号、密码(验证码)成功。第一次调用成功后，后续调用会直接返回第一次的调用结果，不会再产生更新操作。多张票同时出票回填时，保证原子性，只允许全部成功或者全部失败，不能存在部分成功或者失败
type AlitripshipordernotifyAPIRequest struct {
	model.Params
	// 出票入参
	_confirmBookRQ *ShipAgentConfirmBookRq
}

// NewAlitripshipordernotifyRequest 初始化AlitripshipordernotifyAPIRequest对象
func NewAlitripshipordernotifyRequest() *AlitripshipordernotifyAPIRequest {
	return &AlitripshipordernotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripshipordernotifyAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripshipordernotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripshipordernotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmBookRQ is ConfirmBookRQ Setter
// 出票入参
func (r *AlitripshipordernotifyAPIRequest) SetConfirmBookRQ(_confirmBookRQ *ShipAgentConfirmBookRq) error {
	r._confirmBookRQ = _confirmBookRQ
	r.Set("confirm_book_r_q", _confirmBookRQ)
	return nil
}

// GetConfirmBookRQ ConfirmBookRQ Getter
func (r AlitripshipordernotifyAPIRequest) GetConfirmBookRQ() *ShipAgentConfirmBookRq {
	return r._confirmBookRQ
}
