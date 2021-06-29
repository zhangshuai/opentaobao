package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【已废除】11-同步历史所有的po单 API返回值 
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单
*/
type AlibabaTmallgenieScpPlanHistoryPoGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanHistoryPoGetResponse
}

// 【已废除】11-同步历史所有的po单 成功返回结果
type AlibabaTmallgenieScpPlanHistoryPoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_history_po_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象封装
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
