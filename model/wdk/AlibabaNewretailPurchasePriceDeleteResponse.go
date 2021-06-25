package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 商户删除采购价 APIResponse
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价
*/
type AlibabaNewretailPurchasePriceDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNewretailPurchasePriceDeleteResponse `json:"alibaba_newretail_purchase_price_delete_response,omitempty"`
}

type AlibabaNewretailPurchasePriceDeleteResponse struct {

    // 拆单结果对象
    Result   *TopBaseResult `json:"result,omitempty"`

}