package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对推广组进行单独移动溢价 API请求
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountUpdateRequest struct {
    model.Params
    // 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
    adgroupIds   []int64
    // 折扣（折扣值在1-400之间）
    mobileDiscount   int64
    // 昵称
    nick   string
}

// 初始化TaobaoSimbaAdgroupMobilediscountUpdateRequest对象
func NewTaobaoSimbaAdgroupMobilediscountUpdateRequest() *TaobaoSimbaAdgroupMobilediscountUpdateRequest{
    return &TaobaoSimbaAdgroupMobilediscountUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.mobilediscount.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupIds Setter
// 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetAdgroupIds(adgroupIds []int64) error {
    r.adgroupIds = adgroupIds
    r.Set("adgroup_ids", adgroupIds)
    return nil
}

// AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetAdgroupIds() []int64 {
    return r.adgroupIds
}
// MobileDiscount Setter
// 折扣（折扣值在1-400之间）
func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetMobileDiscount(mobileDiscount int64) error {
    r.mobileDiscount = mobileDiscount
    r.Set("mobile_discount", mobileDiscount)
    return nil
}

// MobileDiscount Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetMobileDiscount() int64 {
    return r.mobileDiscount
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetNick() string {
    return r.nick
}
