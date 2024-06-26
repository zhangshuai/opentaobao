package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderValidateAPIRequest 商旅酒店API分销下单前校验 API请求
// alitrip.btrip.hotel.distribution.order.validate
//
// 商旅酒店API分销下单前校验
type AlitripBtripHotelDistributionOrderValidateAPIRequest struct {
	model.Params
	// 下单前校验入参
	_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq
}

// NewAlitripBtripHotelDistributionOrderValidateRequest 初始化AlitripBtripHotelDistributionOrderValidateAPIRequest对象
func NewAlitripBtripHotelDistributionOrderValidateRequest() *AlitripBtripHotelDistributionOrderValidateAPIRequest {
	return &AlitripBtripHotelDistributionOrderValidateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionOrderValidateAPIRequest) Reset() {
	r._paramBtripHotelValidateOrderRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelValidateOrderRq is ParamBtripHotelValidateOrderRq Setter
// 下单前校验入参
func (r *AlitripBtripHotelDistributionOrderValidateAPIRequest) SetParamBtripHotelValidateOrderRq(_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq) error {
	r._paramBtripHotelValidateOrderRq = _paramBtripHotelValidateOrderRq
	r.Set("param_btrip_hotel_validate_order_rq", _paramBtripHotelValidateOrderRq)
	return nil
}

// GetParamBtripHotelValidateOrderRq ParamBtripHotelValidateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderValidateAPIRequest) GetParamBtripHotelValidateOrderRq() *BtripHotelValidateOrderRq {
	return r._paramBtripHotelValidateOrderRq
}

var poolAlitripBtripHotelDistributionOrderValidateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionOrderValidateRequest()
	},
}

// GetAlitripBtripHotelDistributionOrderValidateRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderValidateAPIRequest
func GetAlitripBtripHotelDistributionOrderValidateAPIRequest() *AlitripBtripHotelDistributionOrderValidateAPIRequest {
	return poolAlitripBtripHotelDistributionOrderValidateAPIRequest.Get().(*AlitripBtripHotelDistributionOrderValidateAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionOrderValidateAPIRequest 将 AlitripBtripHotelDistributionOrderValidateAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderValidateAPIRequest(v *AlitripBtripHotelDistributionOrderValidateAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderValidateAPIRequest.Put(v)
}
