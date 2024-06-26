package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmAuditAPIRequest 同情用药审核接口 API请求
// alibaba.alihealth.pw.gm.audit
//
// 同情用药审核接口，提供给合作方审核申请单
type AlibabaAlihealthPwGmAuditAPIRequest struct {
	model.Params
	// 入参
	_body *AuditReq
}

// NewAlibabaAlihealthPwGmAuditRequest 初始化AlibabaAlihealthPwGmAuditAPIRequest对象
func NewAlibabaAlihealthPwGmAuditRequest() *AlibabaAlihealthPwGmAuditAPIRequest {
	return &AlibabaAlihealthPwGmAuditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwGmAuditAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmAuditAPIRequest) SetBody(_body *AuditReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetBody() *AuditReq {
	return r._body
}

var poolAlibabaAlihealthPwGmAuditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwGmAuditRequest()
	},
}

// GetAlibabaAlihealthPwGmAuditRequest 从 sync.Pool 获取 AlibabaAlihealthPwGmAuditAPIRequest
func GetAlibabaAlihealthPwGmAuditAPIRequest() *AlibabaAlihealthPwGmAuditAPIRequest {
	return poolAlibabaAlihealthPwGmAuditAPIRequest.Get().(*AlibabaAlihealthPwGmAuditAPIRequest)
}

// ReleaseAlibabaAlihealthPwGmAuditAPIRequest 将 AlibabaAlihealthPwGmAuditAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwGmAuditAPIRequest(v *AlibabaAlihealthPwGmAuditAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwGmAuditAPIRequest.Put(v)
}
