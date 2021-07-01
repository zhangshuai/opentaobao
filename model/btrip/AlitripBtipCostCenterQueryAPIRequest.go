package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部成本中心 API请求
alitrip.btip.cost.center.query

查询外部成本中心
*/
type AlitripBtipCostCenterQueryAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenCostCenterQueryRq
}

// 初始化AlitripBtipCostCenterQueryAPIRequest对象
func NewAlitripBtipCostCenterQueryRequest() *AlitripBtipCostCenterQueryAPIRequest{
    return &AlitripBtipCostCenterQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtipCostCenterQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.btip.cost.center.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtipCostCenterQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtipCostCenterQueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtipCostCenterQueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
    return r._rq
}