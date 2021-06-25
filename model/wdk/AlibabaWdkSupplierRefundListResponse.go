package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口按供应商拉取退款单 APIResponse
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单
*/
type AlibabaWdkSupplierRefundListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSupplierRefundListResponse `json:"alibaba_wdk_supplier_refund_list_response,omitempty"`
}

type AlibabaWdkSupplierRefundListResponse struct {

    // result
    Result   *OrderSyncRefundListResult `json:"result,omitempty"`

}