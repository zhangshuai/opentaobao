package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店房型与房价查询 API请求
taobao.xhotel.baseinfo.room.get

根据outHid/hid获取酒店的房型和价格信息
*/
type TaobaoXhotelBaseinfoRoomGetRequest struct {
    model.Params
    // 卖家系统中的酒店ID。
    outHid   string
    // 用于标示该酒店发布的渠道信息
    vendor   string
    // 是否需要房价基本信息（false为不需要），默认为需要
    isNeedRatePlan   bool
}

// 初始化TaobaoXhotelBaseinfoRoomGetRequest对象
func NewTaobaoXhotelBaseinfoRoomGetRequest() *TaobaoXhotelBaseinfoRoomGetRequest{
    return &TaobaoXhotelBaseinfoRoomGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBaseinfoRoomGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.baseinfo.room.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBaseinfoRoomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutHid Setter
// 卖家系统中的酒店ID。
func (r *TaobaoXhotelBaseinfoRoomGetRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelBaseinfoRoomGetRequest) GetOutHid() string {
    return r.outHid
}
// Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoXhotelBaseinfoRoomGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelBaseinfoRoomGetRequest) GetVendor() string {
    return r.vendor
}
// IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoRoomGetRequest) SetIsNeedRatePlan(isNeedRatePlan bool) error {
    r.isNeedRatePlan = isNeedRatePlan
    r.Set("is_need_rate_plan", isNeedRatePlan)
    return nil
}

// IsNeedRatePlan Getter
func (r TaobaoXhotelBaseinfoRoomGetRequest) GetIsNeedRatePlan() bool {
    return r.isNeedRatePlan
}