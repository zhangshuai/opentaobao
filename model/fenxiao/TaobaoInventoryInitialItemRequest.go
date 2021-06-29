package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存初始化 API请求
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存
*/
type TaobaoInventoryInitialItemRequest struct {
    model.Params
    // 后端商品id
    scItemId   int64
    // 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
    storeInventorys   string
}

// 初始化TaobaoInventoryInitialItemRequest对象
func NewTaobaoInventoryInitialItemRequest() *TaobaoInventoryInitialItemRequest{
    return &TaobaoInventoryInitialItemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryInitialItemRequest) GetApiMethodName() string {
    return "taobao.inventory.initial.item"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryInitialItemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScItemId Setter
// 后端商品id
func (r *TaobaoInventoryInitialItemRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

// ScItemId Getter
func (r TaobaoInventoryInitialItemRequest) GetScItemId() int64 {
    return r.scItemId
}
// StoreInventorys Setter
// 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
func (r *TaobaoInventoryInitialItemRequest) SetStoreInventorys(storeInventorys string) error {
    r.storeInventorys = storeInventorys
    r.Set("store_inventorys", storeInventorys)
    return nil
}

// StoreInventorys Getter
func (r TaobaoInventoryInitialItemRequest) GetStoreInventorys() string {
    return r.storeInventorys
}
