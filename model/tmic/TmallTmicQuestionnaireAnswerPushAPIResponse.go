package tmic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireAnswerPushAPIResponse 提交单题答案 API返回值
// tmall.tmic.questionnaire.answer.push
//
// 问卷单题回答的提交
type TmallTmicQuestionnaireAnswerPushAPIResponse struct {
	model.CommonResponse
	TmallTmicQuestionnaireAnswerPushAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireAnswerPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTmicQuestionnaireAnswerPushAPIResponseModel).Reset()
}

// TmallTmicQuestionnaireAnswerPushAPIResponseModel is 提交单题答案 成功返回结果
type TmallTmicQuestionnaireAnswerPushAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmic_questionnaire_answer_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示语
	BizErrorInfo string `json:"biz_error_info,omitempty" xml:"biz_error_info,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否调用成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireAnswerPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorInfo = ""
	m.BizErrorCode = ""
	m.BizSuccess = false
}

var poolTmallTmicQuestionnaireAnswerPushAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTmicQuestionnaireAnswerPushAPIResponse)
	},
}

// GetTmallTmicQuestionnaireAnswerPushAPIResponse 从 sync.Pool 获取 TmallTmicQuestionnaireAnswerPushAPIResponse
func GetTmallTmicQuestionnaireAnswerPushAPIResponse() *TmallTmicQuestionnaireAnswerPushAPIResponse {
	return poolTmallTmicQuestionnaireAnswerPushAPIResponse.Get().(*TmallTmicQuestionnaireAnswerPushAPIResponse)
}

// ReleaseTmallTmicQuestionnaireAnswerPushAPIResponse 将 TmallTmicQuestionnaireAnswerPushAPIResponse 保存到 sync.Pool
func ReleaseTmallTmicQuestionnaireAnswerPushAPIResponse(v *TmallTmicQuestionnaireAnswerPushAPIResponse) {
	v.Reset()
	poolTmallTmicQuestionnaireAnswerPushAPIResponse.Put(v)
}
