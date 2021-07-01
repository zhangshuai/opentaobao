package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-等级营销活动查询 API请求
taobao.crm.grademkt.member.detail.query

商家通过该接口查询等级营销活动
*/
type TaobaoCrmGrademktMemberDetailQueryAPIRequest struct {
    model.Params
    // 扩展字段
    _feather   string
    // 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    _parameter   string
}

// 初始化TaobaoCrmGrademktMemberDetailQueryAPIRequest对象
func NewTaobaoCrmGrademktMemberDetailQueryRequest() *TaobaoCrmGrademktMemberDetailQueryAPIRequest{
    return &TaobaoCrmGrademktMemberDetailQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailQueryAPIRequest) SetFeather(_feather string) error {
    r._feather = _feather
    r.Set("feather", _feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetFeather() string {
    return r._feather
}
// Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailQueryAPIRequest) SetParameter(_parameter string) error {
    r._parameter = _parameter
    r.Set("parameter", _parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetParameter() string {
    return r._parameter
}