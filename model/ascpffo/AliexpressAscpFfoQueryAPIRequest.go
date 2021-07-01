package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单查询API API请求
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
type AliexpressAscpFfoQueryAPIRequest struct {
    model.Params
    // dto
    _fulfillmentForwardOrderQuery   *FulfillmentForwardOrderQueryDto
}

// 初始化AliexpressAscpFfoQueryAPIRequest对象
func NewAliexpressAscpFfoQueryRequest() *AliexpressAscpFfoQueryAPIRequest{
    return &AliexpressAscpFfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoQueryAPIRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentForwardOrderQuery Setter
// dto
func (r *AliexpressAscpFfoQueryAPIRequest) SetFulfillmentForwardOrderQuery(_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
    r._fulfillmentForwardOrderQuery = _fulfillmentForwardOrderQuery
    r.Set("fulfillment_forward_order_query", _fulfillmentForwardOrderQuery)
    return nil
}

// FulfillmentForwardOrderQuery Getter
func (r AliexpressAscpFfoQueryAPIRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
    return r._fulfillmentForwardOrderQuery
}