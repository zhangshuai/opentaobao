package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG物流签收状态写接口 API请求
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
type TaobaoNextoneLogisticsSignUpdateRequest struct {
    model.Params
    // 退款编号
    refundId   int64
    // 货物签收状态
    signStatus   int64
}

// 初始化TaobaoNextoneLogisticsSignUpdateRequest对象
func NewTaobaoNextoneLogisticsSignUpdateRequest() *TaobaoNextoneLogisticsSignUpdateRequest{
    return &TaobaoNextoneLogisticsSignUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoNextoneLogisticsSignUpdateRequest) GetApiMethodName() string {
    return "taobao.nextone.logistics.sign.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoNextoneLogisticsSignUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoNextoneLogisticsSignUpdateRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoNextoneLogisticsSignUpdateRequest) GetRefundId() int64 {
    return r.refundId
}
// SignStatus Setter
// 货物签收状态
func (r *TaobaoNextoneLogisticsSignUpdateRequest) SetSignStatus(signStatus int64) error {
    r.signStatus = signStatus
    r.Set("sign_status", signStatus)
    return nil
}

// SignStatus Getter
func (r TaobaoNextoneLogisticsSignUpdateRequest) GetSignStatus() int64 {
    return r.signStatus
}
