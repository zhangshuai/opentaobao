package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品 API请求
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品
*/
type TaobaoBanamadpcItemAddAPIRequest struct {
    model.Params
    // 类目id
    _catId   int64
    // 商品的schema xml
    _xml   string
}

// 初始化TaobaoBanamadpcItemAddAPIRequest对象
func NewTaobaoBanamadpcItemAddRequest() *TaobaoBanamadpcItemAddAPIRequest{
    return &TaobaoBanamadpcItemAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemAddAPIRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemAddAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TaobaoBanamadpcItemAddAPIRequest) GetCatId() int64 {
    return r._catId
}
// Xml Setter
// 商品的schema xml
func (r *TaobaoBanamadpcItemAddAPIRequest) SetXml(_xml string) error {
    r._xml = _xml
    r.Set("xml", _xml)
    return nil
}

// Xml Getter
func (r TaobaoBanamadpcItemAddAPIRequest) GetXml() string {
    return r._xml
}