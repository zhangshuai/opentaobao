package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态更新接口 API请求
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
*/
type TaobaoQimenOrderstatusUpdateAPIRequest struct {
    model.Params
    // 星盘派单号
    _allocationCode   string
    // 淘系子订单号
    _orderCodes   []int64
    // 目标门店的商户中心门店编码
    _storeId   int64
    // 业务类型，（枚举值：FAHUO、ZITI）
    _type   string
    // 订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY
    _status   string
    // 操作人
    _operator   string
    // 事件发生时间
    _actionTime   string
}

// 初始化TaobaoQimenOrderstatusUpdateAPIRequest对象
func NewTaobaoQimenOrderstatusUpdateRequest() *TaobaoQimenOrderstatusUpdateAPIRequest{
    return &TaobaoQimenOrderstatusUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.orderstatus.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AllocationCode Setter
// 星盘派单号
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetAllocationCode(_allocationCode string) error {
    r._allocationCode = _allocationCode
    r.Set("allocation_code", _allocationCode)
    return nil
}

// AllocationCode Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetAllocationCode() string {
    return r._allocationCode
}
// OrderCodes Setter
// 淘系子订单号
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetOrderCodes(_orderCodes []int64) error {
    r._orderCodes = _orderCodes
    r.Set("order_codes", _orderCodes)
    return nil
}

// OrderCodes Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetOrderCodes() []int64 {
    return r._orderCodes
}
// StoreId Setter
// 目标门店的商户中心门店编码
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// Type Setter
// 业务类型，（枚举值：FAHUO、ZITI）
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetType() string {
    return r._type
}
// Status Setter
// 订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetStatus() string {
    return r._status
}
// Operator Setter
// 操作人
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetOperator() string {
    return r._operator
}
// ActionTime Setter
// 事件发生时间
func (r *TaobaoQimenOrderstatusUpdateAPIRequest) SetActionTime(_actionTime string) error {
    r._actionTime = _actionTime
    r.Set("action_time", _actionTime)
    return nil
}

// ActionTime Getter
func (r TaobaoQimenOrderstatusUpdateAPIRequest) GetActionTime() string {
    return r._actionTime
}