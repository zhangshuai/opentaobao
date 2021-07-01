package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商直发-商家仓库存查询服务 API请求
alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query

提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。
*/
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest struct {
    model.Params
    // 商家仓库存查询请求参数
    _merchantInventoryQueryRequest   *MerchantInventoryQuery
}

// 初始化AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest() *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest{
    return &AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantInventoryQueryRequest Setter
// 商家仓库存查询请求参数
func (r *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) SetMerchantInventoryQueryRequest(_merchantInventoryQueryRequest *MerchantInventoryQuery) error {
    r._merchantInventoryQueryRequest = _merchantInventoryQueryRequest
    r.Set("merchant_inventory_query_request", _merchantInventoryQueryRequest)
    return nil
}

// MerchantInventoryQueryRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetMerchantInventoryQueryRequest() *MerchantInventoryQuery {
    return r._merchantInventoryQueryRequest
}