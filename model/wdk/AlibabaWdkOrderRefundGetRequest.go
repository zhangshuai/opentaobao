package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单退款按ID查询 API请求
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情
*/
type AlibabaWdkOrderRefundGetRequest struct {
    model.Params
    // 五道口订单列表（子订单）
    bizOrderIds   []int64
    // 退款订单列表
    refundIds   []int64
    // 渠道来源 3：饿了么 4：盒马
    orderFrom   int64
    // 渠道店id
    shopId   string
    // 经营店id
    storeId   string
}

// 初始化AlibabaWdkOrderRefundGetRequest对象
func NewAlibabaWdkOrderRefundGetRequest() *AlibabaWdkOrderRefundGetRequest{
    return &AlibabaWdkOrderRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderRefundGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderIds Setter
// 五道口订单列表（子订单）
func (r *AlibabaWdkOrderRefundGetRequest) SetBizOrderIds(bizOrderIds []int64) error {
    r.bizOrderIds = bizOrderIds
    r.Set("biz_order_ids", bizOrderIds)
    return nil
}

// BizOrderIds Getter
func (r AlibabaWdkOrderRefundGetRequest) GetBizOrderIds() []int64 {
    return r.bizOrderIds
}
// RefundIds Setter
// 退款订单列表
func (r *AlibabaWdkOrderRefundGetRequest) SetRefundIds(refundIds []int64) error {
    r.refundIds = refundIds
    r.Set("refund_ids", refundIds)
    return nil
}

// RefundIds Getter
func (r AlibabaWdkOrderRefundGetRequest) GetRefundIds() []int64 {
    return r.refundIds
}
// OrderFrom Setter
// 渠道来源 3：饿了么 4：盒马
func (r *AlibabaWdkOrderRefundGetRequest) SetOrderFrom(orderFrom int64) error {
    r.orderFrom = orderFrom
    r.Set("order_from", orderFrom)
    return nil
}

// OrderFrom Getter
func (r AlibabaWdkOrderRefundGetRequest) GetOrderFrom() int64 {
    return r.orderFrom
}
// ShopId Setter
// 渠道店id
func (r *AlibabaWdkOrderRefundGetRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkOrderRefundGetRequest) GetShopId() string {
    return r.shopId
}
// StoreId Setter
// 经营店id
func (r *AlibabaWdkOrderRefundGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkOrderRefundGetRequest) GetStoreId() string {
    return r.storeId
}
