package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备历史数据批量获取 API请求
alibaba.campus.device.historydata.get

设备历史数据批量获取
*/
type AlibabaCampusDeviceHistorydataGetRequest struct {
    model.Params
    // workbench
    workBenchContext   *WorkBenchContext
    // 查询条件
    query   *DeviceHistoryBatchQuery
}

// 初始化AlibabaCampusDeviceHistorydataGetRequest对象
func NewAlibabaCampusDeviceHistorydataGetRequest() *AlibabaCampusDeviceHistorydataGetRequest{
    return &AlibabaCampusDeviceHistorydataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceHistorydataGetRequest) GetApiMethodName() string {
    return "alibaba.campus.device.historydata.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceHistorydataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// workbench
func (r *AlibabaCampusDeviceHistorydataGetRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceHistorydataGetRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Query Setter
// 查询条件
func (r *AlibabaCampusDeviceHistorydataGetRequest) SetQuery(query *DeviceHistoryBatchQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceHistorydataGetRequest) GetQuery() *DeviceHistoryBatchQuery {
    return r.query
}
