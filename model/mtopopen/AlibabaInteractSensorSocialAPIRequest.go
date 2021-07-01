package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
社交组件 API请求
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
type AlibabaInteractSensorSocialAPIRequest struct {
    model.Params
    // 系统自动生成
    _id   string
}

// 初始化AlibabaInteractSensorSocialAPIRequest对象
func NewAlibabaInteractSensorSocialRequest() *AlibabaInteractSensorSocialAPIRequest{
    return &AlibabaInteractSensorSocialAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorSocialAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.social"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorSocialAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorSocialAPIRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaInteractSensorSocialAPIRequest) GetId() string {
    return r._id
}