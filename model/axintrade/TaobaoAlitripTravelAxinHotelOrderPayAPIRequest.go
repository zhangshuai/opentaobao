package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderPayAPIRequest 阿信酒店分销订单支付 API请求
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoAlitripTravelAxinHotelOrderPayAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 采购单号
	_purchaseOrderId int64
}

// NewTaobaoAlitripTravelAxinHotelOrderPayRequest 初始化TaobaoAlitripTravelAxinHotelOrderPayAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderPayRequest() *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderPayAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) Reset() {
	r._distributorTid = 0
	r._purchaseOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单号
func (r *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}

var poolTaobaoAlitripTravelAxinHotelOrderPayAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelOrderPayRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderPayRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderPayAPIRequest
func GetTaobaoAlitripTravelAxinHotelOrderPayAPIRequest() *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelOrderPayAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelOrderPayAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderPayAPIRequest 将 TaobaoAlitripTravelAxinHotelOrderPayAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderPayAPIRequest(v *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderPayAPIRequest.Put(v)
}
