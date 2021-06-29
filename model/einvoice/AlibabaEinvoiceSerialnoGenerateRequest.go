package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取统一开票流水号 API请求
alibaba.einvoice.serialno.generate

erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
*/
type AlibabaEinvoiceSerialnoGenerateRequest struct {
    model.Params
}

// 初始化AlibabaEinvoiceSerialnoGenerateRequest对象
func NewAlibabaEinvoiceSerialnoGenerateRequest() *AlibabaEinvoiceSerialnoGenerateRequest{
    return &AlibabaEinvoiceSerialnoGenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceSerialnoGenerateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.serialno.generate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceSerialnoGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
