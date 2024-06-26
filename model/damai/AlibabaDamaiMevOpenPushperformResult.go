package damai

import (
	"sync"
)

// AlibabaDamaiMevOpenPushperformResult 结构体
type AlibabaDamaiMevOpenPushperformResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDamaiMevOpenPushperformResult = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushperformResult)
	},
}

// GetAlibabaDamaiMevOpenPushperformResult() 从对象池中获取AlibabaDamaiMevOpenPushperformResult
func GetAlibabaDamaiMevOpenPushperformResult() *AlibabaDamaiMevOpenPushperformResult {
	return poolAlibabaDamaiMevOpenPushperformResult.Get().(*AlibabaDamaiMevOpenPushperformResult)
}

// ReleaseAlibabaDamaiMevOpenPushperformResult 释放AlibabaDamaiMevOpenPushperformResult
func ReleaseAlibabaDamaiMevOpenPushperformResult(v *AlibabaDamaiMevOpenPushperformResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = false
	v.Success = false
	poolAlibabaDamaiMevOpenPushperformResult.Put(v)
}
