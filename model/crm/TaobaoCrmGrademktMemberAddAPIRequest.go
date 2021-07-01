package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员吸纳 API请求
taobao.crm.grademkt.member.add

商家通过该接口吸纳线上店铺会员。
*/
type TaobaoCrmGrademktMemberAddAPIRequest struct {
    model.Params
    // 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    _parameter   string
    // 系统属性，json格式
    _feather   string
    // 会员nick
    _buyerNick   string
}

// 初始化TaobaoCrmGrademktMemberAddAPIRequest对象
func NewTaobaoCrmGrademktMemberAddRequest() *TaobaoCrmGrademktMemberAddAPIRequest{
    return &TaobaoCrmGrademktMemberAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberAddAPIRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameter Setter
// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberAddAPIRequest) SetParameter(_parameter string) error {
    r._parameter = _parameter
    r.Set("parameter", _parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberAddAPIRequest) GetParameter() string {
    return r._parameter
}
// Feather Setter
// 系统属性，json格式
func (r *TaobaoCrmGrademktMemberAddAPIRequest) SetFeather(_feather string) error {
    r._feather = _feather
    r.Set("feather", _feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberAddAPIRequest) GetFeather() string {
    return r._feather
}
// BuyerNick Setter
// 会员nick
func (r *TaobaoCrmGrademktMemberAddAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmGrademktMemberAddAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}