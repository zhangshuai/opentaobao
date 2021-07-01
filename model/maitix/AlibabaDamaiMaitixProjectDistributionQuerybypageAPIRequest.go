package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目分页查询项目列表服务 API请求
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务
*/
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest struct {
    model.Params
    // 入参param
    _param   *ProjectPageParam
}

// 初始化AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerybypageRequest() *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querybypage"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参param
func (r *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) SetParam(_param *ProjectPageParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetParam() *ProjectPageParam {
    return r._param
}