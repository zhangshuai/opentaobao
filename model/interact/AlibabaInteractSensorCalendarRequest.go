package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) API请求
alibaba.interact.sensor.calendar

天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
*/
type AlibabaInteractSensorCalendarRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorCalendarRequest对象
func NewAlibabaInteractSensorCalendarRequest() *AlibabaInteractSensorCalendarRequest{
    return &AlibabaInteractSensorCalendarRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorCalendarRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.calendar"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorCalendarRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
