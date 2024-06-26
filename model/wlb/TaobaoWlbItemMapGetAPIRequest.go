package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemMapGetAPIRequest 根据物流宝商品ID查询商品映射关系 API请求
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetAPIRequest struct {
	model.Params
	// 要查询映射关系的物流宝商品id
	_itemId int64
}

// NewTaobaoWlbItemMapGetRequest 初始化TaobaoWlbItemMapGetAPIRequest对象
func NewTaobaoWlbItemMapGetRequest() *TaobaoWlbItemMapGetAPIRequest {
	return &TaobaoWlbItemMapGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbItemMapGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemMapGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.map.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemMapGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbItemMapGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 要查询映射关系的物流宝商品id
func (r *TaobaoWlbItemMapGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbItemMapGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoWlbItemMapGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbItemMapGetRequest()
	},
}

// GetTaobaoWlbItemMapGetRequest 从 sync.Pool 获取 TaobaoWlbItemMapGetAPIRequest
func GetTaobaoWlbItemMapGetAPIRequest() *TaobaoWlbItemMapGetAPIRequest {
	return poolTaobaoWlbItemMapGetAPIRequest.Get().(*TaobaoWlbItemMapGetAPIRequest)
}

// ReleaseTaobaoWlbItemMapGetAPIRequest 将 TaobaoWlbItemMapGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbItemMapGetAPIRequest(v *TaobaoWlbItemMapGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbItemMapGetAPIRequest.Put(v)
}
