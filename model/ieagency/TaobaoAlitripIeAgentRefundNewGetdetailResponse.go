package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询申请单详情(新版) API返回值 
taobao.alitrip.ie.agent.refund.new.getdetail

查询申请单详情
*/
type TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewGetdetailResponse
}

// 查询申请单详情(新版) 成功返回结果
type TaobaoAlitripIeAgentRefundNewGetdetailResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getdetail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *RefundOrderQueryDetailRs `json:"result,omitempty" xml:"result,omitempty"`
}
