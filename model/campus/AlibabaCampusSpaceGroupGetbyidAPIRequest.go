package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分组ID查询相关的空间分组信息 API请求
alibaba.campus.space.group.getbyid

根据分组ID查询相关的空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getById
*/
type AlibabaCampusSpaceGroupGetbyidAPIRequest struct {
    model.Params
    // 用户环境
    _param0   *WorkBenchContext
    // 分组ID
    _param1   int64
}

// 初始化AlibabaCampusSpaceGroupGetbyidAPIRequest对象
func NewAlibabaCampusSpaceGroupGetbyidRequest() *AlibabaCampusSpaceGroupGetbyidAPIRequest{
    return &AlibabaCampusSpaceGroupGetbyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.space.group.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户环境
func (r *AlibabaCampusSpaceGroupGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 分组ID
func (r *AlibabaCampusSpaceGroupGetbyidAPIRequest) SetParam1(_param1 int64) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetParam1() int64 {
    return r._param1
}