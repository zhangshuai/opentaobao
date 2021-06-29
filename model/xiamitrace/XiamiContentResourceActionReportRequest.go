package xiamitrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
曲库开放平台内容行为上报接口 API请求
xiami.content.resource.action.report

合作方对接入的曲库开放内容上报行为日志
*/
type XiamiContentResourceActionReportRequest struct {
    model.Params
    // 资源ID
    resourceId   string
    // 行为数量
    num   int64
    // 资源类型（可枚举）: song(歌曲)
    resourceType   string
    // 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
    action   string
    // 来源id，如歌单id
    fromId   string
    // 1推荐2歌单3标签
    fromType   int64
    // 用户id
    openId   string
    // 用户设备id
    utdid   string
}

// 初始化XiamiContentResourceActionReportRequest对象
func NewXiamiContentResourceActionReportRequest() *XiamiContentResourceActionReportRequest{
    return &XiamiContentResourceActionReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentResourceActionReportRequest) GetApiMethodName() string {
    return "xiami.content.resource.action.report"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentResourceActionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResourceId Setter
// 资源ID
func (r *XiamiContentResourceActionReportRequest) SetResourceId(resourceId string) error {
    r.resourceId = resourceId
    r.Set("resource_id", resourceId)
    return nil
}

// ResourceId Getter
func (r XiamiContentResourceActionReportRequest) GetResourceId() string {
    return r.resourceId
}
// Num Setter
// 行为数量
func (r *XiamiContentResourceActionReportRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r XiamiContentResourceActionReportRequest) GetNum() int64 {
    return r.num
}
// ResourceType Setter
// 资源类型（可枚举）: song(歌曲)
func (r *XiamiContentResourceActionReportRequest) SetResourceType(resourceType string) error {
    r.resourceType = resourceType
    r.Set("resource_type", resourceType)
    return nil
}

// ResourceType Getter
func (r XiamiContentResourceActionReportRequest) GetResourceType() string {
    return r.resourceType
}
// Action Setter
// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
func (r *XiamiContentResourceActionReportRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r XiamiContentResourceActionReportRequest) GetAction() string {
    return r.action
}
// FromId Setter
// 来源id，如歌单id
func (r *XiamiContentResourceActionReportRequest) SetFromId(fromId string) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

// FromId Getter
func (r XiamiContentResourceActionReportRequest) GetFromId() string {
    return r.fromId
}
// FromType Setter
// 1推荐2歌单3标签
func (r *XiamiContentResourceActionReportRequest) SetFromType(fromType int64) error {
    r.fromType = fromType
    r.Set("from_type", fromType)
    return nil
}

// FromType Getter
func (r XiamiContentResourceActionReportRequest) GetFromType() int64 {
    return r.fromType
}
// OpenId Setter
// 用户id
func (r *XiamiContentResourceActionReportRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r XiamiContentResourceActionReportRequest) GetOpenId() string {
    return r.openId
}
// Utdid Setter
// 用户设备id
func (r *XiamiContentResourceActionReportRequest) SetUtdid(utdid string) error {
    r.utdid = utdid
    r.Set("utdid", utdid)
    return nil
}

// Utdid Getter
func (r XiamiContentResourceActionReportRequest) GetUtdid() string {
    return r.utdid
}