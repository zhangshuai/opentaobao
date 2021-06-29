package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品修改规则获取 API请求
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
type TmallItemStoreUpdateSchemaGetRequest struct {
    model.Params
    // 主商品ID
    mainItemId   int64
    // 门店ID
    storeId   int64
}

// 初始化TmallItemStoreUpdateSchemaGetRequest对象
func NewTmallItemStoreUpdateSchemaGetRequest() *TmallItemStoreUpdateSchemaGetRequest{
    return &TmallItemStoreUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemStoreUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.store.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemStoreUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainItemId Setter
// 主商品ID
func (r *TmallItemStoreUpdateSchemaGetRequest) SetMainItemId(mainItemId int64) error {
    r.mainItemId = mainItemId
    r.Set("main_item_id", mainItemId)
    return nil
}

// MainItemId Getter
func (r TmallItemStoreUpdateSchemaGetRequest) GetMainItemId() int64 {
    return r.mainItemId
}
// StoreId Setter
// 门店ID
func (r *TmallItemStoreUpdateSchemaGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TmallItemStoreUpdateSchemaGetRequest) GetStoreId() int64 {
    return r.storeId
}
