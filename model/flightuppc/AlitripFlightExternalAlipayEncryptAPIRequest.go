package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipayencryptAPIRequest 支付宝小程序明文加密 API请求
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
type AlitripflightexternalalipayencryptAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayEncryptReq *AlipayEncryptReq
}

// NewAlitripflightexternalalipayencryptRequest 初始化AlitripflightexternalalipayencryptAPIRequest对象
func NewAlitripflightexternalalipayencryptRequest() *AlitripflightexternalalipayencryptAPIRequest {
	return &AlitripflightexternalalipayencryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightexternalalipayencryptAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightexternalalipayencryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightexternalalipayencryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayEncryptReq is AlipayEncryptReq Setter
// 入参结构体
func (r *AlitripflightexternalalipayencryptAPIRequest) SetAlipayEncryptReq(_alipayEncryptReq *AlipayEncryptReq) error {
	r._alipayEncryptReq = _alipayEncryptReq
	r.Set("alipay_encrypt_req", _alipayEncryptReq)
	return nil
}

// GetAlipayEncryptReq AlipayEncryptReq Getter
func (r AlitripflightexternalalipayencryptAPIRequest) GetAlipayEncryptReq() *AlipayEncryptReq {
	return r._alipayEncryptReq
}
