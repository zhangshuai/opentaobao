package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单申请接口 API返回值 
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单
*/
type TaobaoFuwuSpConfirmApplyAPIResponse struct {
    model.CommonResponse
    TaobaoFuwuSpConfirmApplyResponse
}

// 内购服务确认单申请接口 成功返回结果
type TaobaoFuwuSpConfirmApplyResponse struct {
    XMLName xml.Name `xml:"fuwu_sp_confirm_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的是服务市场的确认单ID
    ApplyResult   int64 `json:"apply_result,omitempty" xml:"apply_result,omitempty"`
}
