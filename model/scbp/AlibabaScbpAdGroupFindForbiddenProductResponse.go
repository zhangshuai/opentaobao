package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽品 APIResponse
alibaba.scbp.ad.group.find.forbidden.product

查询屏蔽品
*/
type AlibabaScbpAdGroupFindForbiddenProductAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdGroupFindForbiddenProductResponse `json:"alibaba_scbp_ad_group_find_forbidden_product_response,omitempty"`
}

type AlibabaScbpAdGroupFindForbiddenProductResponse struct {

    // 返回列表
    ResultList   []ForbiddenProductDto `json:"result_list,omitempty"`

}