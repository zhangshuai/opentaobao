package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsreturnbillgetAPIRequest 获取销退收货信息 API请求
// taobao.wlb.wms.return.bill.get
//
// 通过订单号获取单个销退单收货信息。
type TaobaowlbwmsreturnbillgetAPIRequest struct {
	model.Params
	// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_orderCode string
	// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_cnOrderCode string
}

// NewTaobaowlbwmsreturnbillgetRequest 初始化TaobaowlbwmsreturnbillgetAPIRequest对象
func NewTaobaowlbwmsreturnbillgetRequest() *TaobaowlbwmsreturnbillgetAPIRequest {
	return &TaobaowlbwmsreturnbillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsreturnbillgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.return.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsreturnbillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsreturnbillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaowlbwmsreturnbillgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbwmsreturnbillgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaowlbwmsreturnbillgetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaowlbwmsreturnbillgetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}
