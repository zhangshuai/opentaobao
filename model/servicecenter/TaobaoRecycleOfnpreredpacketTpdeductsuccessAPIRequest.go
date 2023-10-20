package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest 回收商同步前置补贴红包的代扣成功事件 API请求
// taobao.recycle.ofnpreredpacket.tpdeductsuccess
//
// 回收商-&gt;天猫后端，同步前置补贴红包的代扣成功事件
type TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest struct {
	model.Params
	// 变化的金额
	_deductAmount int64
	// 旧机单id
	_oldOrderId int64
}

// NewTaobaorecycleofnpreredpackettpdeductsuccessRequest 初始化TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest对象
func NewTaobaorecycleofnpreredpackettpdeductsuccessRequest() *TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest {
	return &TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnpreredpacket.tpdeductsuccess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeductAmount is DeductAmount Setter
// 变化的金额
func (r *TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) SetDeductAmount(_deductAmount int64) error {
	r._deductAmount = _deductAmount
	r.Set("deduct_amount", _deductAmount)
	return nil
}

// GetDeductAmount DeductAmount Getter
func (r TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) GetDeductAmount() int64 {
	return r._deductAmount
}

// SetOldOrderId is OldOrderId Setter
// 旧机单id
func (r *TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}
