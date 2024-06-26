package alicom

import (
	"sync"
)

// SequenceCalls 结构体
type SequenceCalls struct {
	// 主叫放音
	CallNoPlayCode string `json:"call_no_play_code,omitempty" xml:"call_no_play_code,omitempty"`
	// 被叫号码
	CalledNo string `json:"called_no,omitempty" xml:"called_no,omitempty"`
	// 被叫号显
	CalledDisplayNo string `json:"called_display_no,omitempty" xml:"called_display_no,omitempty"`
	// 被叫放音
	CalledNoPlayCode string `json:"called_no_play_code,omitempty" xml:"called_no_play_code,omitempty"`
	// 摘机后主叫侧的放音编码，多个文件用英文逗号分隔。
	CalledNoCallerPlayCode string `json:"called_no_caller_play_code,omitempty" xml:"called_no_caller_play_code,omitempty"`
	// 顺振序号，从1开始
	PollingNo int64 `json:"polling_no,omitempty" xml:"polling_no,omitempty"`
}

var poolSequenceCalls = sync.Pool{
	New: func() any {
		return new(SequenceCalls)
	},
}

// GetSequenceCalls() 从对象池中获取SequenceCalls
func GetSequenceCalls() *SequenceCalls {
	return poolSequenceCalls.Get().(*SequenceCalls)
}

// ReleaseSequenceCalls 释放SequenceCalls
func ReleaseSequenceCalls(v *SequenceCalls) {
	v.CallNoPlayCode = ""
	v.CalledNo = ""
	v.CalledDisplayNo = ""
	v.CalledNoPlayCode = ""
	v.CalledNoCallerPlayCode = ""
	v.PollingNo = 0
	poolSequenceCalls.Put(v)
}
