package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceinvoiceapplyupdateAPIRequest 商家开票申请单状态回传 API请求
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
type AlibabaeinvoiceinvoiceapplyupdateAPIRequest struct {
	model.Params
	// 申请单id
	_applyId string
	// 扩展信息,目前用于回传文本及物流消息
	_exInfo string
	// 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
	_status int64
}

// NewAlibabaeinvoiceinvoiceapplyupdateRequest 初始化AlibabaeinvoiceinvoiceapplyupdateAPIRequest对象
func NewAlibabaeinvoiceinvoiceapplyupdateRequest() *AlibabaeinvoiceinvoiceapplyupdateAPIRequest {
	return &AlibabaeinvoiceinvoiceapplyupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *AlibabaeinvoiceinvoiceapplyupdateAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetExInfo is ExInfo Setter
// 扩展信息,目前用于回传文本及物流消息
func (r *AlibabaeinvoiceinvoiceapplyupdateAPIRequest) SetExInfo(_exInfo string) error {
	r._exInfo = _exInfo
	r.Set("ex_info", _exInfo)
	return nil
}

// GetExInfo ExInfo Getter
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetExInfo() string {
	return r._exInfo
}

// SetStatus is Status Setter
// 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
func (r *AlibabaeinvoiceinvoiceapplyupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaeinvoiceinvoiceapplyupdateAPIRequest) GetStatus() int64 {
	return r._status
}
