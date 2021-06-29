package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV绑定外部门店id和外部商品id API请求
alibaba.alihealth.dental.item.bind

ISV绑定外部门店id和外部商品id
*/
type AlibabaAlihealthDentalItemBindRequest struct {
    model.Params
    // bind_list
    bindList   []StoreItemRelRequest
    // 类型 1 天猫门店 2 支付宝门店
    type   int64
}

// 初始化AlibabaAlihealthDentalItemBindRequest对象
func NewAlibabaAlihealthDentalItemBindRequest() *AlibabaAlihealthDentalItemBindRequest{
    return &AlibabaAlihealthDentalItemBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemBindRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BindList Setter
// bind_list
func (r *AlibabaAlihealthDentalItemBindRequest) SetBindList(bindList []StoreItemRelRequest) error {
    r.bindList = bindList
    r.Set("bind_list", bindList)
    return nil
}

// BindList Getter
func (r AlibabaAlihealthDentalItemBindRequest) GetBindList() []StoreItemRelRequest {
    return r.bindList
}
// Type Setter
// 类型 1 天猫门店 2 支付宝门店
func (r *AlibabaAlihealthDentalItemBindRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthDentalItemBindRequest) GetType() int64 {
    return r.type
}