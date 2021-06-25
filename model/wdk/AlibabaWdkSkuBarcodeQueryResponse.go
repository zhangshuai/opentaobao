package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品条码查询接口 APIResponse
alibaba.wdk.sku.barcode.query

查询商品编码，支持一品多码
*/
type AlibabaWdkSkuBarcodeQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuBarcodeQueryResponse `json:"alibaba_wdk_sku_barcode_query_response,omitempty"`
}

type AlibabaWdkSkuBarcodeQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuBarcodeQueryApiResults `json:"result,omitempty"`

}