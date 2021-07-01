package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则编辑商品 API请求
tmall.item.schema.update

天猫根据规则编辑商品
*/
type TmallItemSchemaUpdateAPIRequest struct {
    model.Params
    // 需要编辑的商品ID
    _itemId   int64
    // 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
    _categoryId   int64
    // 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
    _productId   int64
    // 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
    _xmlData   string
}

// 初始化TmallItemSchemaUpdateAPIRequest对象
func NewTmallItemSchemaUpdateRequest() *TmallItemSchemaUpdateAPIRequest{
    return &TmallItemSchemaUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSchemaUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.item.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSchemaUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 需要编辑的商品ID
func (r *TmallItemSchemaUpdateAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemSchemaUpdateAPIRequest) GetItemId() int64 {
    return r._itemId
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
func (r *TmallItemSchemaUpdateAPIRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TmallItemSchemaUpdateAPIRequest) GetCategoryId() int64 {
    return r._categoryId
}
// ProductId Setter
// 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
func (r *TmallItemSchemaUpdateAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallItemSchemaUpdateAPIRequest) GetProductId() int64 {
    return r._productId
}
// XmlData Setter
// 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
func (r *TmallItemSchemaUpdateAPIRequest) SetXmlData(_xmlData string) error {
    r._xmlData = _xmlData
    r.Set("xml_data", _xmlData)
    return nil
}

// XmlData Getter
func (r TmallItemSchemaUpdateAPIRequest) GetXmlData() string {
    return r._xmlData
}