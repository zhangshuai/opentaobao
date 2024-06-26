package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceProdApplyGetAPIRequest 查询发票申请 API请求
// alibaba.einvoice.prod.apply.get
//
// 查询申请的详细信息，包含申请所关联的发票摘要信息+板式文件+预览图；
// 场景使用：1、业务前台收到申请状态变更消息后，调用此接口查询申请详情；2、主动补偿查询：当指定了自动开票，且发票申请长时间未收到状态变更通知时，可能存在丢消息的情况，此时可主动查询该申请，然后更新本地工单状态。
type AlibabaEinvoiceProdApplyGetAPIRequest struct {
	model.Params
	// 查询申请请求
	_invoiceApplyQueryDto *InvoiceApplyDtlQueryDto
}

// NewAlibabaEinvoiceProdApplyGetRequest 初始化AlibabaEinvoiceProdApplyGetAPIRequest对象
func NewAlibabaEinvoiceProdApplyGetRequest() *AlibabaEinvoiceProdApplyGetAPIRequest {
	return &AlibabaEinvoiceProdApplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceProdApplyGetAPIRequest) Reset() {
	r._invoiceApplyQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceProdApplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.prod.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceProdApplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceProdApplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceApplyQueryDto is InvoiceApplyQueryDto Setter
// 查询申请请求
func (r *AlibabaEinvoiceProdApplyGetAPIRequest) SetInvoiceApplyQueryDto(_invoiceApplyQueryDto *InvoiceApplyDtlQueryDto) error {
	r._invoiceApplyQueryDto = _invoiceApplyQueryDto
	r.Set("invoice_apply_query_dto", _invoiceApplyQueryDto)
	return nil
}

// GetInvoiceApplyQueryDto InvoiceApplyQueryDto Getter
func (r AlibabaEinvoiceProdApplyGetAPIRequest) GetInvoiceApplyQueryDto() *InvoiceApplyDtlQueryDto {
	return r._invoiceApplyQueryDto
}

var poolAlibabaEinvoiceProdApplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceProdApplyGetRequest()
	},
}

// GetAlibabaEinvoiceProdApplyGetRequest 从 sync.Pool 获取 AlibabaEinvoiceProdApplyGetAPIRequest
func GetAlibabaEinvoiceProdApplyGetAPIRequest() *AlibabaEinvoiceProdApplyGetAPIRequest {
	return poolAlibabaEinvoiceProdApplyGetAPIRequest.Get().(*AlibabaEinvoiceProdApplyGetAPIRequest)
}

// ReleaseAlibabaEinvoiceProdApplyGetAPIRequest 将 AlibabaEinvoiceProdApplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceProdApplyGetAPIRequest(v *AlibabaEinvoiceProdApplyGetAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceProdApplyGetAPIRequest.Put(v)
}
