package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rate删除接口 API请求
taobao.xhotel.rate.delete

酒店产品库rate删除
*/
type TaobaoXhotelRateDeleteRequest struct {
    model.Params
    // 系统商，一般不用填写，使用须申请
    vendor   string
    // 商家价格政策编码
    rateplanCode   string
    // 商家房型ID
    outRid   string
}

// 初始化TaobaoXhotelRateDeleteRequest对象
func NewTaobaoXhotelRateDeleteRequest() *TaobaoXhotelRateDeleteRequest{
    return &TaobaoXhotelRateDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateDeleteRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateDeleteRequest) GetVendor() string {
    return r.vendor
}
// RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateDeleteRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateDeleteRequest) GetRateplanCode() string {
    return r.rateplanCode
}
// OutRid Setter
// 商家房型ID
func (r *TaobaoXhotelRateDeleteRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRateDeleteRequest) GetOutRid() string {
    return r.outRid
}