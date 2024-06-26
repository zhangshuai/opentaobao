package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductMapAddAPIRequest 创建分销和后端商品映射关系 API请求
// taobao.fenxiao.product.map.add
//
// 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddAPIRequest struct {
	model.Params
	// 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
	_skuIds string
	// 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
	_scItemIds string
	// 分销产品id。
	_productId int64
	// 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
	_scItemId int64
	// 是否需要校验商家编码，true不校验，false校验。
	_notCheckOuterCode bool
}

// NewTaobaoFenxiaoProductMapAddRequest 初始化TaobaoFenxiaoProductMapAddAPIRequest对象
func NewTaobaoFenxiaoProductMapAddRequest() *TaobaoFenxiaoProductMapAddAPIRequest {
	return &TaobaoFenxiaoProductMapAddAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductMapAddAPIRequest) Reset() {
	r._skuIds = ""
	r._scItemIds = ""
	r._productId = 0
	r._scItemId = 0
	r._notCheckOuterCode = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.map.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuIds is SkuIds Setter
// 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
func (r *TaobaoFenxiaoProductMapAddAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetScItemIds is ScItemIds Setter
// 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
func (r *TaobaoFenxiaoProductMapAddAPIRequest) SetScItemIds(_scItemIds string) error {
	r._scItemIds = _scItemIds
	r.Set("sc_item_ids", _scItemIds)
	return nil
}

// GetScItemIds ScItemIds Getter
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetScItemIds() string {
	return r._scItemIds
}

// SetProductId is ProductId Setter
// 分销产品id。
func (r *TaobaoFenxiaoProductMapAddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetScItemId is ScItemId Setter
// 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
func (r *TaobaoFenxiaoProductMapAddAPIRequest) SetScItemId(_scItemId int64) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetScItemId() int64 {
	return r._scItemId
}

// SetNotCheckOuterCode is NotCheckOuterCode Setter
// 是否需要校验商家编码，true不校验，false校验。
func (r *TaobaoFenxiaoProductMapAddAPIRequest) SetNotCheckOuterCode(_notCheckOuterCode bool) error {
	r._notCheckOuterCode = _notCheckOuterCode
	r.Set("not_check_outer_code", _notCheckOuterCode)
	return nil
}

// GetNotCheckOuterCode NotCheckOuterCode Getter
func (r TaobaoFenxiaoProductMapAddAPIRequest) GetNotCheckOuterCode() bool {
	return r._notCheckOuterCode
}

var poolTaobaoFenxiaoProductMapAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductMapAddRequest()
	},
}

// GetTaobaoFenxiaoProductMapAddRequest 从 sync.Pool 获取 TaobaoFenxiaoProductMapAddAPIRequest
func GetTaobaoFenxiaoProductMapAddAPIRequest() *TaobaoFenxiaoProductMapAddAPIRequest {
	return poolTaobaoFenxiaoProductMapAddAPIRequest.Get().(*TaobaoFenxiaoProductMapAddAPIRequest)
}

// ReleaseTaobaoFenxiaoProductMapAddAPIRequest 将 TaobaoFenxiaoProductMapAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductMapAddAPIRequest(v *TaobaoFenxiaoProductMapAddAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductMapAddAPIRequest.Put(v)
}
