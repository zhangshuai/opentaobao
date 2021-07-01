package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中心仓查网格仓 API请求
wdk.logistic.network.warehouse.resource.relation.query.from

盒马集市，中心仓查询网格仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest struct {
    model.Params
    // 查询参数
    _paramPageQueryWarehouseResourceRelationByFromRequest   *PageQueryWarehouseResourceRelationByFromRequest
}

// 初始化WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest对象
func NewWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest{
    return &WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.resource.relation.query.from"
}

// IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPageQueryWarehouseResourceRelationByFromRequest Setter
// 查询参数
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) SetParamPageQueryWarehouseResourceRelationByFromRequest(_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest) error {
    r._paramPageQueryWarehouseResourceRelationByFromRequest = _paramPageQueryWarehouseResourceRelationByFromRequest
    r.Set("param_page_query_warehouse_resource_relation_by_from_request", _paramPageQueryWarehouseResourceRelationByFromRequest)
    return nil
}

// ParamPageQueryWarehouseResourceRelationByFromRequest Getter
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetParamPageQueryWarehouseResourceRelationByFromRequest() *PageQueryWarehouseResourceRelationByFromRequest {
    return r._paramPageQueryWarehouseResourceRelationByFromRequest
}