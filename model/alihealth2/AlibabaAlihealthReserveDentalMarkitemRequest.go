package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标记商品是否可预约 API请求
alibaba.alihealth.reserve.dental.markitem

标记商品是否可预约
*/
type AlibabaAlihealthReserveDentalMarkitemRequest struct {
    model.Params
    // 平台商品id
    itemId   int64
    // 是否可预约，1.可预约 0.不可预约
    status   int64
}

// 初始化AlibabaAlihealthReserveDentalMarkitemRequest对象
func NewAlibabaAlihealthReserveDentalMarkitemRequest() *AlibabaAlihealthReserveDentalMarkitemRequest{
    return &AlibabaAlihealthReserveDentalMarkitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalMarkitemRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.markitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalMarkitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 平台商品id
func (r *AlibabaAlihealthReserveDentalMarkitemRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaAlihealthReserveDentalMarkitemRequest) GetItemId() int64 {
    return r.itemId
}
// Status Setter
// 是否可预约，1.可预约 0.不可预约
func (r *AlibabaAlihealthReserveDentalMarkitemRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthReserveDentalMarkitemRequest) GetStatus() int64 {
    return r.status
}