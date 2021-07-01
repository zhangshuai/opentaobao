package ticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票规则信息查询接口 API返回值 
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
type AlitripTicketRuleQueryAPIResponse struct {
    model.CommonResponse
    AlitripTicketRuleQueryAPIResponseModel
}

// 【门票API2.0】门票规则信息查询接口 成功返回结果
type AlitripTicketRuleQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_ticket_rule_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 门票规则信息
    FirstResult   *TicketRuleParam `json:"first_result,omitempty" xml:"first_result,omitempty"`
}