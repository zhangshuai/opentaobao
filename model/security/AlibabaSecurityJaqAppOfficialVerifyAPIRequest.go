package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappofficialverifyAPIRequest 聚安全验证官方应用接口 API请求
// alibaba.security.jaq.app.official.verify
//
// 接入用户来查询应用是否为官方应用
type AlibabasecurityjaqappofficialverifyAPIRequest struct {
	model.Params
	// 验证参数
	_officialAppVerifyRequest *OfficialAppVerifyRequest
}

// NewAlibabasecurityjaqappofficialverifyRequest 初始化AlibabasecurityjaqappofficialverifyAPIRequest对象
func NewAlibabasecurityjaqappofficialverifyRequest() *AlibabasecurityjaqappofficialverifyAPIRequest {
	return &AlibabasecurityjaqappofficialverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqappofficialverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.official.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqappofficialverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqappofficialverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfficialAppVerifyRequest is OfficialAppVerifyRequest Setter
// 验证参数
func (r *AlibabasecurityjaqappofficialverifyAPIRequest) SetOfficialAppVerifyRequest(_officialAppVerifyRequest *OfficialAppVerifyRequest) error {
	r._officialAppVerifyRequest = _officialAppVerifyRequest
	r.Set("official_app_verify_request", _officialAppVerifyRequest)
	return nil
}

// GetOfficialAppVerifyRequest OfficialAppVerifyRequest Getter
func (r AlibabasecurityjaqappofficialverifyAPIRequest) GetOfficialAppVerifyRequest() *OfficialAppVerifyRequest {
	return r._officialAppVerifyRequest
}
