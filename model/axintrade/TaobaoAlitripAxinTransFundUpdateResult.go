package axintrade

import (
	"sync"
)

// TaobaoAlitripAxinTransFundUpdateResult 结构体
type TaobaoAlitripAxinTransFundUpdateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 提示信息
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 成功数量
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripAxinTransFundUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransFundUpdateResult)
	},
}

// GetTaobaoAlitripAxinTransFundUpdateResult() 从对象池中获取TaobaoAlitripAxinTransFundUpdateResult
func GetTaobaoAlitripAxinTransFundUpdateResult() *TaobaoAlitripAxinTransFundUpdateResult {
	return poolTaobaoAlitripAxinTransFundUpdateResult.Get().(*TaobaoAlitripAxinTransFundUpdateResult)
}

// ReleaseTaobaoAlitripAxinTransFundUpdateResult 释放TaobaoAlitripAxinTransFundUpdateResult
func ReleaseTaobaoAlitripAxinTransFundUpdateResult(v *TaobaoAlitripAxinTransFundUpdateResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.InfoMsg = ""
	v.Data = 0
	v.NeedRetry = false
	v.Success = false
	poolTaobaoAlitripAxinTransFundUpdateResult.Put(v)
}
