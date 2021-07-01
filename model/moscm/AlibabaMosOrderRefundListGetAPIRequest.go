package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询交易退货信息 API请求
alibaba.mos.order.refund.list.get

批量查询多个退货单的退货明细
*/
type AlibabaMosOrderRefundListGetAPIRequest struct {
    model.Params
    // 退换货查询条件
    _rmaCriteria   *RmaCriteria
    // 分页信息
    _paginator   *Paginator
}

// 初始化AlibabaMosOrderRefundListGetAPIRequest对象
func NewAlibabaMosOrderRefundListGetRequest() *AlibabaMosOrderRefundListGetAPIRequest{
    return &AlibabaMosOrderRefundListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderRefundListGetAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.order.refund.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderRefundListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RmaCriteria Setter
// 退换货查询条件
func (r *AlibabaMosOrderRefundListGetAPIRequest) SetRmaCriteria(_rmaCriteria *RmaCriteria) error {
    r._rmaCriteria = _rmaCriteria
    r.Set("rma_criteria", _rmaCriteria)
    return nil
}

// RmaCriteria Getter
func (r AlibabaMosOrderRefundListGetAPIRequest) GetRmaCriteria() *RmaCriteria {
    return r._rmaCriteria
}
// Paginator Setter
// 分页信息
func (r *AlibabaMosOrderRefundListGetAPIRequest) SetPaginator(_paginator *Paginator) error {
    r._paginator = _paginator
    r.Set("paginator", _paginator)
    return nil
}

// Paginator Getter
func (r AlibabaMosOrderRefundListGetAPIRequest) GetPaginator() *Paginator {
    return r._paginator
}