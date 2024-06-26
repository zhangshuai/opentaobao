package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallDeviceItemPromotionGetAPIRequest 智能硬件上商品优惠获取 API请求
// tmall.device.item.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TmallDeviceItemPromotionGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallDeviceItemPromotionGetRequest 初始化TmallDeviceItemPromotionGetAPIRequest对象
func NewTmallDeviceItemPromotionGetRequest() *TmallDeviceItemPromotionGetAPIRequest {
	return &TmallDeviceItemPromotionGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallDeviceItemPromotionGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallDeviceItemPromotionGetAPIRequest) GetApiMethodName() string {
	return "tmall.device.item.promotion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallDeviceItemPromotionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallDeviceItemPromotionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallDeviceItemPromotionGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallDeviceItemPromotionGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallDeviceItemPromotionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallDeviceItemPromotionGetRequest()
	},
}

// GetTmallDeviceItemPromotionGetRequest 从 sync.Pool 获取 TmallDeviceItemPromotionGetAPIRequest
func GetTmallDeviceItemPromotionGetAPIRequest() *TmallDeviceItemPromotionGetAPIRequest {
	return poolTmallDeviceItemPromotionGetAPIRequest.Get().(*TmallDeviceItemPromotionGetAPIRequest)
}

// ReleaseTmallDeviceItemPromotionGetAPIRequest 将 TmallDeviceItemPromotionGetAPIRequest 放入 sync.Pool
func ReleaseTmallDeviceItemPromotionGetAPIRequest(v *TmallDeviceItemPromotionGetAPIRequest) {
	v.Reset()
	poolTmallDeviceItemPromotionGetAPIRequest.Put(v)
}
