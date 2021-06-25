package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
应用风险概要信息查询接口 APIResponse
alibaba.security.jaq.app.risksummary.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
*/
type AlibabaSecurityJaqAppRisksummaryGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqAppRisksummaryGetResponse `json:"alibaba_security_jaq_app_risksummary_get_response,omitempty"`
}

type AlibabaSecurityJaqAppRisksummaryGetResponse struct {

    // 应用扫描概要信息
    Result   *RiskSummary `json:"result,omitempty"`

}