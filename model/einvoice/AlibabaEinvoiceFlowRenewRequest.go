package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单(入驻、加盘、续约)续约 API请求
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放
*/
type AlibabaEinvoiceFlowRenewRequest struct {
    model.Params
    // 续约请求参数
    invoiceFlowRenewDto   *InvoiceFlowRenewDto
}

// 初始化AlibabaEinvoiceFlowRenewRequest对象
func NewAlibabaEinvoiceFlowRenewRequest() *AlibabaEinvoiceFlowRenewRequest{
    return &AlibabaEinvoiceFlowRenewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRenewRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.renew"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRenewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceFlowRenewDto Setter
// 续约请求参数
func (r *AlibabaEinvoiceFlowRenewRequest) SetInvoiceFlowRenewDto(invoiceFlowRenewDto *InvoiceFlowRenewDto) error {
    r.invoiceFlowRenewDto = invoiceFlowRenewDto
    r.Set("invoice_flow_renew_dto", invoiceFlowRenewDto)
    return nil
}

// InvoiceFlowRenewDto Getter
func (r AlibabaEinvoiceFlowRenewRequest) GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
    return r.invoiceFlowRenewDto
}
