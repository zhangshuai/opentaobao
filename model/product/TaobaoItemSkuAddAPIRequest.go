package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加SKU API请求
taobao.item.sku.add

新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户
*/
type TaobaoItemSkuAddAPIRequest struct {
    model.Params
    // Sku所属商品数字id。必选
    _numIid   int64
    // Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
    _properties   string
    // Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
    _quantity   int64
    // Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
    _price   float64
    // Sku的商家外部id
    _outerId   string
    // sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
    _itemPrice   float64
    // Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    _lang   string
    // 忽略警告提示.
    _ignorewarning   string
}

// 初始化TaobaoItemSkuAddAPIRequest对象
func NewTaobaoItemSkuAddRequest() *TaobaoItemSkuAddAPIRequest{
    return &TaobaoItemSkuAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuAddAPIRequest) GetApiMethodName() string {
    return "taobao.item.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// Sku所属商品数字id。必选
func (r *TaobaoItemSkuAddAPIRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemSkuAddAPIRequest) GetNumIid() int64 {
    return r._numIid
}
// Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
func (r *TaobaoItemSkuAddAPIRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoItemSkuAddAPIRequest) GetProperties() string {
    return r._properties
}
// Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
func (r *TaobaoItemSkuAddAPIRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoItemSkuAddAPIRequest) GetQuantity() int64 {
    return r._quantity
}
// Price Setter
// Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemSkuAddAPIRequest) SetPrice(_price float64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoItemSkuAddAPIRequest) GetPrice() float64 {
    return r._price
}
// OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuAddAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemSkuAddAPIRequest) GetOuterId() string {
    return r._outerId
}
// ItemPrice Setter
// sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
func (r *TaobaoItemSkuAddAPIRequest) SetItemPrice(_itemPrice float64) error {
    r._itemPrice = _itemPrice
    r.Set("item_price", _itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoItemSkuAddAPIRequest) GetItemPrice() float64 {
    return r._itemPrice
}
// Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuAddAPIRequest) SetLang(_lang string) error {
    r._lang = _lang
    r.Set("lang", _lang)
    return nil
}

// Lang Getter
func (r TaobaoItemSkuAddAPIRequest) GetLang() string {
    return r._lang
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuAddAPIRequest) SetIgnorewarning(_ignorewarning string) error {
    r._ignorewarning = _ignorewarning
    r.Set("ignorewarning", _ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemSkuAddAPIRequest) GetIgnorewarning() string {
    return r._ignorewarning
}