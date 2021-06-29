package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐藏titleBar API请求
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
type AlibabaInteractSensorTitlebarhideRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorTitlebarhideRequest对象
func NewAlibabaInteractSensorTitlebarhideRequest() *AlibabaInteractSensorTitlebarhideRequest{
    return &AlibabaInteractSensorTitlebarhideRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTitlebarhideRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.titlebarhide"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTitlebarhideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
