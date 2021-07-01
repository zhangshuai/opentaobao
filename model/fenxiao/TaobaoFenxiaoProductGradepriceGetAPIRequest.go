package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
等级折扣查询 API请求
taobao.fenxiao.product.gradeprice.get

等级折扣查询
*/
type TaobaoFenxiaoProductGradepriceGetAPIRequest struct {
    model.Params
    // 产品id
    _productId   int64
    // skuId
    _skuId   int64
    // 经、代销模式（1：代销、2：经销）
    _tradeType   int64
}

// 初始化TaobaoFenxiaoProductGradepriceGetAPIRequest对象
func NewTaobaoFenxiaoProductGradepriceGetRequest() *TaobaoFenxiaoProductGradepriceGetAPIRequest{
    return &TaobaoFenxiaoProductGradepriceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.gradeprice.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品id
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetProductId() int64 {
    return r._productId
}
// SkuId Setter
// skuId
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetSkuId() int64 {
    return r._skuId
}
// TradeType Setter
// 经、代销模式（1：代销、2：经销）
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetTradeType(_tradeType int64) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetTradeType() int64 {
    return r._tradeType
}