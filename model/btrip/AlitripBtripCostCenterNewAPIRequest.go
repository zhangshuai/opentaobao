package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建外部成本中心 API请求
alitrip.btrip.cost.center.new

新建外部成本中心
*/
type AlitripBtripCostCenterNewAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenCostCenterSaveRq
}

// 初始化AlitripBtripCostCenterNewAPIRequest对象
func NewAlitripBtripCostCenterNewRequest() *AlitripBtripCostCenterNewAPIRequest{
    return &AlitripBtripCostCenterNewAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterNewAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.new"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterNewAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterNewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterNewAPIRequest) GetRq() *OpenCostCenterSaveRq {
    return r._rq
}