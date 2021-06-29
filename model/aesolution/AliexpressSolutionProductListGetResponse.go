package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Get product list API返回值 
aliexpress.solution.product.list.get

Get product list
*/
type AliexpressSolutionProductListGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductListGetResponse
}

// Get product list 成功返回结果
type AliexpressSolutionProductListGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_list_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ItemListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}