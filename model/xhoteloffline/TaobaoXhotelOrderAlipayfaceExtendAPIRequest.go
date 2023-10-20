package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderalipayfaceextendAPIRequest 信用住订单延住接口 API请求
// taobao.xhotel.order.alipayface.extend
//
// 信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
type TaobaoxhotelorderalipayfaceextendAPIRequest struct {
	model.Params
	// 延住后的离店日期(最多总共入住天数不能超过9间夜)
	_checkOut string
	// 卖家系统订单号,和tid必须至少传入一个
	_outOrderId string
	// 延住日期段的每日房价信息,注意不包括原有的日期房价
	_extendDailyPriceInfo string
	// 阿里旅行订单id,和outOrderId必须至少传入一个
	_tid int64
	// 延住房费,注意不包含原有的房费金额,单位为分
	_extendFee int64
}

// NewTaobaoxhotelorderalipayfaceextendRequest 初始化TaobaoxhotelorderalipayfaceextendAPIRequest对象
func NewTaobaoxhotelorderalipayfaceextendRequest() *TaobaoxhotelorderalipayfaceextendAPIRequest {
	return &TaobaoxhotelorderalipayfaceextendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.extend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckOut is CheckOut Setter
// 延住后的离店日期(最多总共入住天数不能超过9间夜)
func (r *TaobaoxhotelorderalipayfaceextendAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetOutOrderId is OutOrderId Setter
// 卖家系统订单号,和tid必须至少传入一个
func (r *TaobaoxhotelorderalipayfaceextendAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetExtendDailyPriceInfo is ExtendDailyPriceInfo Setter
// 延住日期段的每日房价信息,注意不包括原有的日期房价
func (r *TaobaoxhotelorderalipayfaceextendAPIRequest) SetExtendDailyPriceInfo(_extendDailyPriceInfo string) error {
	r._extendDailyPriceInfo = _extendDailyPriceInfo
	r.Set("extend_daily_price_info", _extendDailyPriceInfo)
	return nil
}

// GetExtendDailyPriceInfo ExtendDailyPriceInfo Getter
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetExtendDailyPriceInfo() string {
	return r._extendDailyPriceInfo
}

// SetTid is Tid Setter
// 阿里旅行订单id,和outOrderId必须至少传入一个
func (r *TaobaoxhotelorderalipayfaceextendAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetTid() int64 {
	return r._tid
}

// SetExtendFee is ExtendFee Setter
// 延住房费,注意不包含原有的房费金额,单位为分
func (r *TaobaoxhotelorderalipayfaceextendAPIRequest) SetExtendFee(_extendFee int64) error {
	r._extendFee = _extendFee
	r.Set("extend_fee", _extendFee)
	return nil
}

// GetExtendFee ExtendFee Getter
func (r TaobaoxhotelorderalipayfaceextendAPIRequest) GetExtendFee() int64 {
	return r._extendFee
}
