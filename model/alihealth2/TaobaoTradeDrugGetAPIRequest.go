package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugGetAPIRequest 查询商家未确认订单列表 API请求
// taobao.trade.drug.get
//
// 可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
type TaobaoTradeDrugGetAPIRequest struct {
	model.Params
	// 店铺id
	_storeId int64
	// 返回记录数，超过20按20条返回数据
	_maxSize int64
	// true-商家下所有店铺的待确认订单, false—指定店铺的订单
	_isAll bool
}

// NewTaobaoTradeDrugGetRequest 初始化TaobaoTradeDrugGetAPIRequest对象
func NewTaobaoTradeDrugGetRequest() *TaobaoTradeDrugGetAPIRequest {
	return &TaobaoTradeDrugGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeDrugGetAPIRequest) Reset() {
	r._storeId = 0
	r._maxSize = 0
	r._isAll = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeDrugGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 店铺id
func (r *TaobaoTradeDrugGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoTradeDrugGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetMaxSize is MaxSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaoTradeDrugGetAPIRequest) SetMaxSize(_maxSize int64) error {
	r._maxSize = _maxSize
	r.Set("max_size", _maxSize)
	return nil
}

// GetMaxSize MaxSize Getter
func (r TaobaoTradeDrugGetAPIRequest) GetMaxSize() int64 {
	return r._maxSize
}

// SetIsAll is IsAll Setter
// true-商家下所有店铺的待确认订单, false—指定店铺的订单
func (r *TaobaoTradeDrugGetAPIRequest) SetIsAll(_isAll bool) error {
	r._isAll = _isAll
	r.Set("is_all", _isAll)
	return nil
}

// GetIsAll IsAll Getter
func (r TaobaoTradeDrugGetAPIRequest) GetIsAll() bool {
	return r._isAll
}

var poolTaobaoTradeDrugGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeDrugGetRequest()
	},
}

// GetTaobaoTradeDrugGetRequest 从 sync.Pool 获取 TaobaoTradeDrugGetAPIRequest
func GetTaobaoTradeDrugGetAPIRequest() *TaobaoTradeDrugGetAPIRequest {
	return poolTaobaoTradeDrugGetAPIRequest.Get().(*TaobaoTradeDrugGetAPIRequest)
}

// ReleaseTaobaoTradeDrugGetAPIRequest 将 TaobaoTradeDrugGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeDrugGetAPIRequest(v *TaobaoTradeDrugGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeDrugGetAPIRequest.Put(v)
}
