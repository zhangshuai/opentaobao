package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取鉴权接口 API请求
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权
*/
type AlibabaInteractCouponApplyAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractCouponApplyAPIRequest对象
func NewAlibabaInteractCouponApplyRequest() *AlibabaInteractCouponApplyAPIRequest{
    return &AlibabaInteractCouponApplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractCouponApplyAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.coupon.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractCouponApplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}