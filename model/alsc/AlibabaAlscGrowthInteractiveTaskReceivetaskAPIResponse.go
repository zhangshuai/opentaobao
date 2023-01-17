package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse 领取任务 API返回值
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
type AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel
}

// AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel is 领取任务 成功返回结果
type AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_receivetask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
