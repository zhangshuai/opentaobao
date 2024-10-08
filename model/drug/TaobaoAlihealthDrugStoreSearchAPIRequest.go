package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugStoreSearchAPIRequest 药品店内搜索 API请求
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
type TaobaoAlihealthDrugStoreSearchAPIRequest struct {
	model.Params
	// 搜索关键字
	_keyword string
	// 店铺ID
	_shopId string
	// 每页显示数量
	_pageSize int64
	// 页码
	_pageNo int64
}

// NewTaobaoAlihealthDrugStoreSearchRequest 初始化TaobaoAlihealthDrugStoreSearchAPIRequest对象
func NewTaobaoAlihealthDrugStoreSearchRequest() *TaobaoAlihealthDrugStoreSearchAPIRequest {
	return &TaobaoAlihealthDrugStoreSearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) Reset() {
	r._keyword = ""
	r._shopId = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.store.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetShopId is ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetShopId() string {
	return r._shopId
}

// SetPageSize is PageSize Setter
// 每页显示数量
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoAlihealthDrugStoreSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAlihealthDrugStoreSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoAlihealthDrugStoreSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlihealthDrugStoreSearchRequest()
	},
}

// GetTaobaoAlihealthDrugStoreSearchRequest 从 sync.Pool 获取 TaobaoAlihealthDrugStoreSearchAPIRequest
func GetTaobaoAlihealthDrugStoreSearchAPIRequest() *TaobaoAlihealthDrugStoreSearchAPIRequest {
	return poolTaobaoAlihealthDrugStoreSearchAPIRequest.Get().(*TaobaoAlihealthDrugStoreSearchAPIRequest)
}

// ReleaseTaobaoAlihealthDrugStoreSearchAPIRequest 将 TaobaoAlihealthDrugStoreSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlihealthDrugStoreSearchAPIRequest(v *TaobaoAlihealthDrugStoreSearchAPIRequest) {
	v.Reset()
	poolTaobaoAlihealthDrugStoreSearchAPIRequest.Put(v)
}
