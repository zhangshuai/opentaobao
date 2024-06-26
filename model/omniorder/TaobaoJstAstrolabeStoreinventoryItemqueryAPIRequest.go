package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest 库存查询接口 API请求
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest struct {
	model.Params
	// 门店信息
	_stores []Store
}

// NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest 初始化TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest() *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) Reset() {
	r._stores = r._stores[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店信息
func (r *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetStores() []Store {
	return r._stores
}

var poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest()
	},
}

// GetTaobaoJstAstrolabeStoreinventoryItemqueryRequest 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest
func GetTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest() *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest {
	return poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest.Get().(*TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest 将 TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest(v *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest.Put(v)
}
