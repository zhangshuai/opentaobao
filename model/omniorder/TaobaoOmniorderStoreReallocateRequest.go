package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rellocate API请求
taobao.omniorder.store.reallocate

门店发货提供改派接口
*/
type TaobaoOmniorderStoreReallocateRequest struct {
    model.Params
    // 主订单号
    mainOrderId   int64
    // 子订单号
    subOrderIds   []int64
    // 门店Id
    storeId   int64
    // 电商仓code
    warehouseCode   string
}

// 初始化TaobaoOmniorderStoreReallocateRequest对象
func NewTaobaoOmniorderStoreReallocateRequest() *TaobaoOmniorderStoreReallocateRequest{
    return &TaobaoOmniorderStoreReallocateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreReallocateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.reallocate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreReallocateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单号
func (r *TaobaoOmniorderStoreReallocateRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniorderStoreReallocateRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// SubOrderIds Setter
// 子订单号
func (r *TaobaoOmniorderStoreReallocateRequest) SetSubOrderIds(subOrderIds []int64) error {
    r.subOrderIds = subOrderIds
    r.Set("sub_order_ids", subOrderIds)
    return nil
}

// SubOrderIds Getter
func (r TaobaoOmniorderStoreReallocateRequest) GetSubOrderIds() []int64 {
    return r.subOrderIds
}
// StoreId Setter
// 门店Id
func (r *TaobaoOmniorderStoreReallocateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreReallocateRequest) GetStoreId() int64 {
    return r.storeId
}
// WarehouseCode Setter
// 电商仓code
func (r *TaobaoOmniorderStoreReallocateRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoOmniorderStoreReallocateRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
