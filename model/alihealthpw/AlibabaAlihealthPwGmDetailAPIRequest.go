package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmDetailAPIRequest 同情用药申请单详情接口 API请求
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
type AlibabaAlihealthPwGmDetailAPIRequest struct {
	model.Params
	// 入参
	_body *DetailForBreq
}

// NewAlibabaAlihealthPwGmDetailRequest 初始化AlibabaAlihealthPwGmDetailAPIRequest对象
func NewAlibabaAlihealthPwGmDetailRequest() *AlibabaAlihealthPwGmDetailAPIRequest {
	return &AlibabaAlihealthPwGmDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmDetailAPIRequest) SetBody(_body *DetailForBreq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetBody() *DetailForBreq {
	return r._body
}
