package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询取消履约的原因列表 API请求
tmall.nr.fulfill.cancel.reason.query

新零售门店业务查询取消上门揽件的原因列表
*/
type TmallNrFulfillCancelReasonQueryRequest struct {
    model.Params
    // 商家的sellerID
    sellerId   int64
    // 淘宝交易的主订单号
    mainOrderId   int64
}

// 初始化TmallNrFulfillCancelReasonQueryRequest对象
func NewTmallNrFulfillCancelReasonQueryRequest() *TmallNrFulfillCancelReasonQueryRequest{
    return &TmallNrFulfillCancelReasonQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelReasonQueryRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.cancel.reason.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelReasonQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 商家的sellerID
func (r *TmallNrFulfillCancelReasonQueryRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallNrFulfillCancelReasonQueryRequest) GetSellerId() int64 {
    return r.sellerId
}
// MainOrderId Setter
// 淘宝交易的主订单号
func (r *TmallNrFulfillCancelReasonQueryRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TmallNrFulfillCancelReasonQueryRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}