package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发票申请 API请求
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
*/
type AlibabaEinvoiceProdApplyAPIRequest struct {
    model.Params
    // 申请开票请求
    _paramInvoiceApplyDto   *InvoiceApplyDto
}

// 初始化AlibabaEinvoiceProdApplyAPIRequest对象
func NewAlibabaEinvoiceProdApplyRequest() *AlibabaEinvoiceProdApplyAPIRequest{
    return &AlibabaEinvoiceProdApplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceProdApplyAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.prod.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceProdApplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamInvoiceApplyDto Setter
// 申请开票请求
func (r *AlibabaEinvoiceProdApplyAPIRequest) SetParamInvoiceApplyDto(_paramInvoiceApplyDto *InvoiceApplyDto) error {
    r._paramInvoiceApplyDto = _paramInvoiceApplyDto
    r.Set("param_invoice_apply_dto", _paramInvoiceApplyDto)
    return nil
}

// ParamInvoiceApplyDto Getter
func (r AlibabaEinvoiceProdApplyAPIRequest) GetParamInvoiceApplyDto() *InvoiceApplyDto {
    return r._paramInvoiceApplyDto
}