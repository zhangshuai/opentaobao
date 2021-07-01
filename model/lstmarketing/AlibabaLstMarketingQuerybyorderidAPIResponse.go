package lstmarketing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询营销信息 API返回值 
alibaba.lst.marketing.querybyorderid

根据订单查询营销信息
*/
type AlibabaLstMarketingQuerybyorderidAPIResponse struct {
    model.CommonResponse
    AlibabaLstMarketingQuerybyorderidAPIResponseModel
}

// 根据订单查询营销信息 成功返回结果
type AlibabaLstMarketingQuerybyorderidAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lst_marketing_querybyorderid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *AlibabaLstMarketingQuerybyorderidResultDto `json:"result,omitempty" xml:"result,omitempty"`
}