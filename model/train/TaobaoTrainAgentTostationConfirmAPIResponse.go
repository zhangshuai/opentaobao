package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTostationConfirmAPIResponse 线下票确认送票至车站服务 API返回值
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
type TaobaoTrainAgentTostationConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentTostationConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTostationConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentTostationConfirmAPIResponseModel).Reset()
}

// TaobaoTrainAgentTostationConfirmAPIResponseModel is 线下票确认送票至车站服务 成功返回结果
type TaobaoTrainAgentTostationConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_tostation_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorMsgCode string `json:"error_msg_code,omitempty" xml:"error_msg_code,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTostationConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsgCode = ""
	m.ExtendParams = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentTostationConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentTostationConfirmAPIResponse)
	},
}

// GetTaobaoTrainAgentTostationConfirmAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentTostationConfirmAPIResponse
func GetTaobaoTrainAgentTostationConfirmAPIResponse() *TaobaoTrainAgentTostationConfirmAPIResponse {
	return poolTaobaoTrainAgentTostationConfirmAPIResponse.Get().(*TaobaoTrainAgentTostationConfirmAPIResponse)
}

// ReleaseTaobaoTrainAgentTostationConfirmAPIResponse 将 TaobaoTrainAgentTostationConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentTostationConfirmAPIResponse(v *TaobaoTrainAgentTostationConfirmAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentTostationConfirmAPIResponse.Put(v)
}
