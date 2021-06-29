package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 API返回值 
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenCombineitemQueryResponse
}

// 组合货品关系查询接口 成功返回结果
type TaobaoQimenCombineitemQueryResponse struct {
    XMLName xml.Name `xml:"qimen_combineitem_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
