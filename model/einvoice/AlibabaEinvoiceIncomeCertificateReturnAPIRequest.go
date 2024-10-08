package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeCertificateReturnAPIRequest 服务商回传进项认证结果 API请求
// alibaba.einvoice.income.certificate.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
type AlibabaEinvoiceIncomeCertificateReturnAPIRequest struct {
	model.Params
	// 错误码，success=false时必填
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// 请求标识
	_reqIndex string
	// 认证步骤，1=勾选，2=汇总，3=确认
	_step int64
	// 认证结果，true=成功，false=失败
	_success bool
}

// NewAlibabaEinvoiceIncomeCertificateReturnRequest 初始化AlibabaEinvoiceIncomeCertificateReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeCertificateReturnRequest() *AlibabaEinvoiceIncomeCertificateReturnAPIRequest {
	return &AlibabaEinvoiceIncomeCertificateReturnAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) Reset() {
	r._errorCode = ""
	r._errorMessage = ""
	r._reqIndex = ""
	r._step = 0
	r._success = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.certificate.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetReqIndex is ReqIndex Setter
// 请求标识
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetStep is Step Setter
// 认证步骤，1=勾选，2=汇总，3=确认
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) SetStep(_step int64) error {
	r._step = _step
	r.Set("step", _step)
	return nil
}

// GetStep Step Getter
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetStep() int64 {
	return r._step
}

// SetSuccess is Success Setter
// 认证结果，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaEinvoiceIncomeCertificateReturnAPIRequest) GetSuccess() bool {
	return r._success
}

var poolAlibabaEinvoiceIncomeCertificateReturnAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceIncomeCertificateReturnRequest()
	},
}

// GetAlibabaEinvoiceIncomeCertificateReturnRequest 从 sync.Pool 获取 AlibabaEinvoiceIncomeCertificateReturnAPIRequest
func GetAlibabaEinvoiceIncomeCertificateReturnAPIRequest() *AlibabaEinvoiceIncomeCertificateReturnAPIRequest {
	return poolAlibabaEinvoiceIncomeCertificateReturnAPIRequest.Get().(*AlibabaEinvoiceIncomeCertificateReturnAPIRequest)
}

// ReleaseAlibabaEinvoiceIncomeCertificateReturnAPIRequest 将 AlibabaEinvoiceIncomeCertificateReturnAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceIncomeCertificateReturnAPIRequest(v *AlibabaEinvoiceIncomeCertificateReturnAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceIncomeCertificateReturnAPIRequest.Put(v)
}
