package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemMapQueryAPIRequest 查找IC商品或分销商品与后端商品的关联信息 API请求
// taobao.scitem.map.query
//
// 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryAPIRequest struct {
	model.Params
	// map_type为1：前台ic商品id<br/>map_type为2：分销productid
	_itemId int64
	// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
	_skuId int64
}

// NewTaobaoScitemMapQueryRequest 初始化TaobaoScitemMapQueryAPIRequest对象
func NewTaobaoScitemMapQueryRequest() *TaobaoScitemMapQueryAPIRequest {
	return &TaobaoScitemMapQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoScitemMapQueryAPIRequest) Reset() {
	r._itemId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapQueryAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoScitemMapQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// map_type为1：前台ic商品id&lt;br/&gt;map_type为2：分销productid
func (r *TaobaoScitemMapQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoScitemMapQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// map_type为1：前台ic商品skuid &lt;br/&gt;map_type为2：分销商品的skuid
func (r *TaobaoScitemMapQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoScitemMapQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoScitemMapQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoScitemMapQueryRequest()
	},
}

// GetTaobaoScitemMapQueryRequest 从 sync.Pool 获取 TaobaoScitemMapQueryAPIRequest
func GetTaobaoScitemMapQueryAPIRequest() *TaobaoScitemMapQueryAPIRequest {
	return poolTaobaoScitemMapQueryAPIRequest.Get().(*TaobaoScitemMapQueryAPIRequest)
}

// ReleaseTaobaoScitemMapQueryAPIRequest 将 TaobaoScitemMapQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoScitemMapQueryAPIRequest(v *TaobaoScitemMapQueryAPIRequest) {
	v.Reset()
	poolTaobaoScitemMapQueryAPIRequest.Put(v)
}
