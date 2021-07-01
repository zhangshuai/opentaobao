package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收商同步前置补贴红包的代扣成功事件 API请求
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件
*/
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest struct {
    model.Params
    // 变化的金额
    _deductAmount   int64
    // 旧机单id
    _oldOrderId   int64
}

// 初始化TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest对象
func NewTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest() *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest{
    return &TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetApiMethodName() string {
    return "taobao.recycle.ofnpreredpacket.tpdeductsuccess"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeductAmount Setter
// 变化的金额
func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) SetDeductAmount(_deductAmount int64) error {
    r._deductAmount = _deductAmount
    r.Set("deduct_amount", _deductAmount)
    return nil
}

// DeductAmount Getter
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetDeductAmount() int64 {
    return r._deductAmount
}
// OldOrderId Setter
// 旧机单id
func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) SetOldOrderId(_oldOrderId int64) error {
    r._oldOrderId = _oldOrderId
    r.Set("old_order_id", _oldOrderId)
    return nil
}

// OldOrderId Getter
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetOldOrderId() int64 {
    return r._oldOrderId
}