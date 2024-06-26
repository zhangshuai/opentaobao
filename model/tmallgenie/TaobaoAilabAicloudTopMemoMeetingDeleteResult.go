package tmallgenie

import (
	"sync"
)

// TaobaoAilabAicloudTopMemoMeetingDeleteResult 结构体
type TaobaoAilabAicloudTopMemoMeetingDeleteResult struct {
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 服务的实际返回结果
	Meeting *Meeting `json:"meeting,omitempty" xml:"meeting,omitempty"`
}

var poolTaobaoAilabAicloudTopMemoMeetingDeleteResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoMeetingDeleteResult)
	},
}

// GetTaobaoAilabAicloudTopMemoMeetingDeleteResult() 从对象池中获取TaobaoAilabAicloudTopMemoMeetingDeleteResult
func GetTaobaoAilabAicloudTopMemoMeetingDeleteResult() *TaobaoAilabAicloudTopMemoMeetingDeleteResult {
	return poolTaobaoAilabAicloudTopMemoMeetingDeleteResult.Get().(*TaobaoAilabAicloudTopMemoMeetingDeleteResult)
}

// ReleaseTaobaoAilabAicloudTopMemoMeetingDeleteResult 释放TaobaoAilabAicloudTopMemoMeetingDeleteResult
func ReleaseTaobaoAilabAicloudTopMemoMeetingDeleteResult(v *TaobaoAilabAicloudTopMemoMeetingDeleteResult) {
	v.Message = ""
	v.StatusCode = 0
	v.Meeting = nil
	poolTaobaoAilabAicloudTopMemoMeetingDeleteResult.Put(v)
}
