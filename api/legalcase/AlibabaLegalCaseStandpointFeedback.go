package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointFeedback 新增或更新 反馈口径(采纳口径/不采纳口径)
// alibaba.legal.case.standpoint.feedback
//
// 新增或更新 反馈口径(采纳口径/不采纳口径)
func AlibabaLegalCaseStandpointFeedback(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointFeedbackAPIRequest, resp *legalcase.AlibabaLegalCaseStandpointFeedbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
