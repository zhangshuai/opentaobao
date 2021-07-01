package globalvirtual

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际虚拟商品发码服务 API请求
alibaba.global.virtual.sendcode

global virtual send code service
*/
type AlibabaGlobalVirtualSendcodeAPIRequest struct {
    model.Params
    // trade order id
    _tradeOrderLineId   int64
    // code list
    _codeList   []VirtualCertificateDo
}

// 初始化AlibabaGlobalVirtualSendcodeAPIRequest对象
func NewAlibabaGlobalVirtualSendcodeRequest() *AlibabaGlobalVirtualSendcodeAPIRequest{
    return &AlibabaGlobalVirtualSendcodeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiMethodName() string {
    return "alibaba.global.virtual.sendcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeOrderLineId Setter
// trade order id
func (r *AlibabaGlobalVirtualSendcodeAPIRequest) SetTradeOrderLineId(_tradeOrderLineId int64) error {
    r._tradeOrderLineId = _tradeOrderLineId
    r.Set("trade_order_line_id", _tradeOrderLineId)
    return nil
}

// TradeOrderLineId Getter
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetTradeOrderLineId() int64 {
    return r._tradeOrderLineId
}
// CodeList Setter
// code list
func (r *AlibabaGlobalVirtualSendcodeAPIRequest) SetCodeList(_codeList []VirtualCertificateDo) error {
    r._codeList = _codeList
    r.Set("code_list", _codeList)
    return nil
}

// CodeList Getter
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetCodeList() []VirtualCertificateDo {
    return r._codeList
}