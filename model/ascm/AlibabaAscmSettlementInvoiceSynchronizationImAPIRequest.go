package ascm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
英迈发票同步到结算 API请求
alibaba.ascm.settlement.invoice.synchronization.im

外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
*/
type AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest struct {
    model.Params
    // im invoice xml
    _xmlDataSlot   string
}

// 初始化AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest对象
func NewAlibabaAscmSettlementInvoiceSynchronizationImRequest() *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest{
    return &AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetApiMethodName() string {
    return "alibaba.ascm.settlement.invoice.synchronization.im"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// XmlDataSlot Setter
// im invoice xml
func (r *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) SetXmlDataSlot(_xmlDataSlot string) error {
    r._xmlDataSlot = _xmlDataSlot
    r.Set("xml_data_slot", _xmlDataSlot)
    return nil
}

// XmlDataSlot Getter
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetXmlDataSlot() string {
    return r._xmlDataSlot
}