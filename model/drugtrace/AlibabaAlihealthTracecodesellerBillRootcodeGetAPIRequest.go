package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest 获取最外层包装码 API请求
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
type AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest struct {
	model.Params
	// 用户身份认证
	_appCode string
	// 码
	_code string
}

// NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest 初始化AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest对象
func NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest() *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest {
	return &AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) Reset() {
	r._appCode = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.rootcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 用户身份认证
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetCode is Code Setter
// 码
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerBillRootcodeGetRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest
func GetAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest() *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest {
	return poolAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest.Get().(*AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest 将 AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest(v *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest.Put(v)
}
