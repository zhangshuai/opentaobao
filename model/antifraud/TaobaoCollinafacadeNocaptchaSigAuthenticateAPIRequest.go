package antifraud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocollinafacadenocaptchasigauthenticateAPIRequest 人机识别 API请求
// taobao.collinafacade.nocaptcha.sig.authenticate
//
// 人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
type TaobaocollinafacadenocaptchasigauthenticateAPIRequest struct {
	model.Params
	// 签名串校验接口入参
	_sigAuthenticateContext *SigAuthenticateContext
}

// NewTaobaocollinafacadenocaptchasigauthenticateRequest 初始化TaobaocollinafacadenocaptchasigauthenticateAPIRequest对象
func NewTaobaocollinafacadenocaptchasigauthenticateRequest() *TaobaocollinafacadenocaptchasigauthenticateAPIRequest {
	return &TaobaocollinafacadenocaptchasigauthenticateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocollinafacadenocaptchasigauthenticateAPIRequest) GetApiMethodName() string {
	return "taobao.collinafacade.nocaptcha.sig.authenticate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocollinafacadenocaptchasigauthenticateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocollinafacadenocaptchasigauthenticateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSigAuthenticateContext is SigAuthenticateContext Setter
// 签名串校验接口入参
func (r *TaobaocollinafacadenocaptchasigauthenticateAPIRequest) SetSigAuthenticateContext(_sigAuthenticateContext *SigAuthenticateContext) error {
	r._sigAuthenticateContext = _sigAuthenticateContext
	r.Set("sig_authenticate_context", _sigAuthenticateContext)
	return nil
}

// GetSigAuthenticateContext SigAuthenticateContext Getter
func (r TaobaocollinafacadenocaptchasigauthenticateAPIRequest) GetSigAuthenticateContext() *SigAuthenticateContext {
	return r._sigAuthenticateContext
}
