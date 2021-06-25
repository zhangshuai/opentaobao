package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退款打款 APIResponse
alibaba.wdk.reverse.refund

五道口退款打款开放能力接口
*/
type AlibabaWdkReverseRefundAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkReverseRefundResponse `json:"alibaba_wdk_reverse_refund_response,omitempty"`
}

type AlibabaWdkReverseRefundResponse struct {

    // 接口返回model
    Result   *AlibabaWdkReverseRefundResult `json:"result,omitempty"`

}