package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分补录 APIResponse
alibaba.alsc.crm.point.extracharge

积分补录
*/
type AlibabaAlscCrmPointExtrachargeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmPointExtrachargeResponse `json:"alibaba_alsc_crm_point_extracharge_response,omitempty"`
}

type AlibabaAlscCrmPointExtrachargeResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}