package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版创建 APIResponse
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版
*/
type AlibabaWdkCouponTemplateCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponTemplateCreateResponse `json:"alibaba_wdk_coupon_template_create_response,omitempty"`
}

type AlibabaWdkCouponTemplateCreateResponse struct {

    // 结果
    Result   *AlibabaWdkCouponTemplateCreateApiResult `json:"result,omitempty"`

}