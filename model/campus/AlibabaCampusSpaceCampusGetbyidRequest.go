package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区id获取园区信息 API请求
alibaba.campus.space.campus.getbyid

根据园区id获取园区信息
HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
HSF方法名称：getCampusById
*/
type AlibabaCampusSpaceCampusGetbyidRequest struct {
    model.Params
    // 园区ID
    param0   *WorkBenchContext
    // 园区ID
    param1   int64
}

// 初始化AlibabaCampusSpaceCampusGetbyidRequest对象
func NewAlibabaCampusSpaceCampusGetbyidRequest() *AlibabaCampusSpaceCampusGetbyidRequest{
    return &AlibabaCampusSpaceCampusGetbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceCampusGetbyidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.campus.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceCampusGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceCampusGetbyidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceCampusGetbyidRequest) GetParam1() int64 {
    return r.param1
}
