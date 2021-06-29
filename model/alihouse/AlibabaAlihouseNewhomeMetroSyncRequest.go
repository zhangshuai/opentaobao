package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地铁数据同步 API请求
alibaba.alihouse.newhome.metro.sync

地铁数据同步
*/
type AlibabaAlihouseNewhomeMetroSyncRequest struct {
    model.Params
    // 地铁入参数据
    baseMetroLineDto   *BaseMetroLineDto
}

// 初始化AlibabaAlihouseNewhomeMetroSyncRequest对象
func NewAlibabaAlihouseNewhomeMetroSyncRequest() *AlibabaAlihouseNewhomeMetroSyncRequest{
    return &AlibabaAlihouseNewhomeMetroSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.metro.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseMetroLineDto Setter
// 地铁入参数据
func (r *AlibabaAlihouseNewhomeMetroSyncRequest) SetBaseMetroLineDto(baseMetroLineDto *BaseMetroLineDto) error {
    r.baseMetroLineDto = baseMetroLineDto
    r.Set("base_metro_line_dto", baseMetroLineDto)
    return nil
}

// BaseMetroLineDto Getter
func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetBaseMetroLineDto() *BaseMetroLineDto {
    return r.baseMetroLineDto
}
