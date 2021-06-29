package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
alibaba.campus.device.openapi.getsimpledevicelist

查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
*/
type AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest struct {
    model.Params
    // 请求发送端信息
    workBenchContext   *WorkBenchContext
    // 多条件查询对象
    query   *DeviceApiQuery
}

// 初始化AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest对象
func NewAlibabaCampusDeviceOpenapiGetsimpledevicelistRequest() *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest{
    return &AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getsimpledevicelist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Query Setter
// 多条件查询对象
func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) SetQuery(query *DeviceApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetQuery() *DeviceApiQuery {
    return r.query
}