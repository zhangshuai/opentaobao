package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取直连酒店（客栈）签约上线进度信息 API请求
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息
*/
type TaobaoXhotelOrderHotelsignQueryAPIRequest struct {
    model.Params
    // 请求流水
    _outUuid   string
    // 商家酒店编码
    _hotelCode   string
    // 商家vendor
    _vendor   string
    // 1
    _type   string
    // 页码
    _pageNo   int64
}

// 初始化TaobaoXhotelOrderHotelsignQueryAPIRequest对象
func NewTaobaoXhotelOrderHotelsignQueryRequest() *TaobaoXhotelOrderHotelsignQueryAPIRequest{
    return &TaobaoXhotelOrderHotelsignQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.hotelsign.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutUuid Setter
// 请求流水
func (r *TaobaoXhotelOrderHotelsignQueryAPIRequest) SetOutUuid(_outUuid string) error {
    r._outUuid = _outUuid
    r.Set("out_uuid", _outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetOutUuid() string {
    return r._outUuid
}
// HotelCode Setter
// 商家酒店编码
func (r *TaobaoXhotelOrderHotelsignQueryAPIRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetHotelCode() string {
    return r._hotelCode
}
// Vendor Setter
// 商家vendor
func (r *TaobaoXhotelOrderHotelsignQueryAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetVendor() string {
    return r._vendor
}
// Type Setter
// 1
func (r *TaobaoXhotelOrderHotelsignQueryAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetType() string {
    return r._type
}
// PageNo Setter
// 页码
func (r *TaobaoXhotelOrderHotelsignQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoXhotelOrderHotelsignQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}