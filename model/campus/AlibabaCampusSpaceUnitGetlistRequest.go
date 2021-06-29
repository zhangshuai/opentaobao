package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询空间单元信息 API请求
alibaba.campus.space.unit.getlist

多条件查询空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getList
*/
type AlibabaCampusSpaceUnitGetlistRequest struct {
    model.Params
    // 查询条件封装
    param0   *WorkBenchContext
    // 查询参数封装
    param1   *SpaceUnitQuery
}

// 初始化AlibabaCampusSpaceUnitGetlistRequest对象
func NewAlibabaCampusSpaceUnitGetlistRequest() *AlibabaCampusSpaceUnitGetlistRequest{
    return &AlibabaCampusSpaceUnitGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceUnitGetlistRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceUnitGetlistRequest) SetParam1(param1 *SpaceUnitQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistRequest) GetParam1() *SpaceUnitQuery {
    return r.param1
}