package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务列表获取接口 API请求
taobao.vmarket.eticket.tasks.get

外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
*/
type TaobaoVmarketEticketTasksGetRequest struct {
    model.Params
    // 卖家家ID(信任卖家不必传，码商可选)
    sellerId   int64
    // 返回结果类型:<br/>1:返回通知失败的订单<br/>2.返回通知成功回调失败的订单
    type   int64
    // 页码。取值范围:大于零的整数; 默认值:1
    pageNo   int64
    // 每页获取条数。默认值40，最小值1，最大值100。
    pageSize   int64
    // 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
    codemerchantId   int64
}

// 初始化TaobaoVmarketEticketTasksGetRequest对象
func NewTaobaoVmarketEticketTasksGetRequest() *TaobaoVmarketEticketTasksGetRequest{
    return &TaobaoVmarketEticketTasksGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketTasksGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.tasks.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketTasksGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 卖家家ID(信任卖家不必传，码商可选)
func (r *TaobaoVmarketEticketTasksGetRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoVmarketEticketTasksGetRequest) GetSellerId() int64 {
    return r.sellerId
}
// Type Setter
// 返回结果类型:<br/>1:返回通知失败的订单<br/>2.返回通知成功回调失败的订单
func (r *TaobaoVmarketEticketTasksGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoVmarketEticketTasksGetRequest) GetType() int64 {
    return r.type
}
// PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaoVmarketEticketTasksGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoVmarketEticketTasksGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值100。
func (r *TaobaoVmarketEticketTasksGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoVmarketEticketTasksGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketTasksGetRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketTasksGetRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}
