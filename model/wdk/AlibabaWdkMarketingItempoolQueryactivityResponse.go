package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查找商品池活动 APIResponse
alibaba.wdk.marketing.itempool.queryactivity

查找商品池活动
*/
type AlibabaWdkMarketingItempoolQueryactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolQueryactivityResponse `json:"alibaba_wdk_marketing_itempool_queryactivity_response,omitempty"`
}

type AlibabaWdkMarketingItempoolQueryactivityResponse struct {

    // 查询返回结果
    Result   *MarketResult `json:"result,omitempty"`

}