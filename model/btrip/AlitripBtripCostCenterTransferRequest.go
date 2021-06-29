package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.cost.center.transfer

商旅成本中心转换为外部成本中心
*/
type AlitripBtripCostCenterTransferRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterTransferRq
}

// 初始化AlitripBtripCostCenterTransferRequest对象
func NewAlitripBtripCostCenterTransferRequest() *AlitripBtripCostCenterTransferRequest{
    return &AlitripBtripCostCenterTransferRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterTransferRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.transfer"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterTransferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterTransferRequest) SetRq(rq *OpenCostCenterTransferRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterTransferRequest) GetRq() *OpenCostCenterTransferRq {
    return r.rq
}