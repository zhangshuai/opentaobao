package idle

import (
	"sync"
)

// AlibabaIdleConsignmentOrderGetResult 结构体
type AlibabaIdleConsignmentOrderGetResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 帮卖订单详情
	Module *ConsignmentOrderTo `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleConsignmentOrderGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentOrderGetResult)
	},
}

// GetAlibabaIdleConsignmentOrderGetResult() 从对象池中获取AlibabaIdleConsignmentOrderGetResult
func GetAlibabaIdleConsignmentOrderGetResult() *AlibabaIdleConsignmentOrderGetResult {
	return poolAlibabaIdleConsignmentOrderGetResult.Get().(*AlibabaIdleConsignmentOrderGetResult)
}

// ReleaseAlibabaIdleConsignmentOrderGetResult 释放AlibabaIdleConsignmentOrderGetResult
func ReleaseAlibabaIdleConsignmentOrderGetResult(v *AlibabaIdleConsignmentOrderGetResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleConsignmentOrderGetResult.Put(v)
}
