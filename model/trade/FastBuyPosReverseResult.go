package trade

import (
	"sync"
)

// FastBuyPosReverseResult 结构体
type FastBuyPosReverseResult struct {
	// 返回的错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回的错误信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 退款状态: 1为成功,2为失败
	ResultStatus int64 `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 调用接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFastBuyPosReverseResult = sync.Pool{
	New: func() any {
		return new(FastBuyPosReverseResult)
	},
}

// GetFastBuyPosReverseResult() 从对象池中获取FastBuyPosReverseResult
func GetFastBuyPosReverseResult() *FastBuyPosReverseResult {
	return poolFastBuyPosReverseResult.Get().(*FastBuyPosReverseResult)
}

// ReleaseFastBuyPosReverseResult 释放FastBuyPosReverseResult
func ReleaseFastBuyPosReverseResult(v *FastBuyPosReverseResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.ResultStatus = 0
	v.Success = false
	poolFastBuyPosReverseResult.Put(v)
}
