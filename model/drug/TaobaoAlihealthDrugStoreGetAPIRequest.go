package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugStoreGetAPIRequest 根据店铺id获取店铺详情 API请求
// taobao.alihealth.drug.store.get
//
// 根据店铺id获取店铺详情
type TaobaoAlihealthDrugStoreGetAPIRequest struct {
	model.Params
	// 店铺ID
	_shopId int64
}

// NewTaobaoAlihealthDrugStoreGetRequest 初始化TaobaoAlihealthDrugStoreGetAPIRequest对象
func NewTaobaoAlihealthDrugStoreGetRequest() *TaobaoAlihealthDrugStoreGetAPIRequest {
	return &TaobaoAlihealthDrugStoreGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugStoreGetAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugStoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlihealthDrugStoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreGetAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoAlihealthDrugStoreGetAPIRequest) GetShopId() int64 {
	return r._shopId
}
