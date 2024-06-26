package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherStatusChangeAPIRequest 优惠券状态更改 API请求
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
type AlibabaAlscCrmVoucherStatusChangeAPIRequest struct {
	model.Params
	// 参数
	_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq
}

// NewAlibabaAlscCrmVoucherStatusChangeRequest 初始化AlibabaAlscCrmVoucherStatusChangeAPIRequest对象
func NewAlibabaAlscCrmVoucherStatusChangeRequest() *AlibabaAlscCrmVoucherStatusChangeAPIRequest {
	return &AlibabaAlscCrmVoucherStatusChangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmVoucherStatusChangeAPIRequest) Reset() {
	r._paramVoucherStatusChangeOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVoucherStatusChangeOpenReq is ParamVoucherStatusChangeOpenReq Setter
// 参数
func (r *AlibabaAlscCrmVoucherStatusChangeAPIRequest) SetParamVoucherStatusChangeOpenReq(_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq) error {
	r._paramVoucherStatusChangeOpenReq = _paramVoucherStatusChangeOpenReq
	r.Set("param_voucher_status_change_open_req", _paramVoucherStatusChangeOpenReq)
	return nil
}

// GetParamVoucherStatusChangeOpenReq ParamVoucherStatusChangeOpenReq Getter
func (r AlibabaAlscCrmVoucherStatusChangeAPIRequest) GetParamVoucherStatusChangeOpenReq() *VoucherStatusChangeOpenReq {
	return r._paramVoucherStatusChangeOpenReq
}

var poolAlibabaAlscCrmVoucherStatusChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmVoucherStatusChangeRequest()
	},
}

// GetAlibabaAlscCrmVoucherStatusChangeRequest 从 sync.Pool 获取 AlibabaAlscCrmVoucherStatusChangeAPIRequest
func GetAlibabaAlscCrmVoucherStatusChangeAPIRequest() *AlibabaAlscCrmVoucherStatusChangeAPIRequest {
	return poolAlibabaAlscCrmVoucherStatusChangeAPIRequest.Get().(*AlibabaAlscCrmVoucherStatusChangeAPIRequest)
}

// ReleaseAlibabaAlscCrmVoucherStatusChangeAPIRequest 将 AlibabaAlscCrmVoucherStatusChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmVoucherStatusChangeAPIRequest(v *AlibabaAlscCrmVoucherStatusChangeAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmVoucherStatusChangeAPIRequest.Put(v)
}
