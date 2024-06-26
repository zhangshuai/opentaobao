package btrip

import (
	"sync"
)

// OpenAccountRs 结构体
type OpenAccountRs struct {
	// 账期结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 账期开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// json数据下载链接，通过HttpClient 获取 并GBK格式解析，链接五分钟有效期
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolOpenAccountRs = sync.Pool{
	New: func() any {
		return new(OpenAccountRs)
	},
}

// GetOpenAccountRs() 从对象池中获取OpenAccountRs
func GetOpenAccountRs() *OpenAccountRs {
	return poolOpenAccountRs.Get().(*OpenAccountRs)
}

// ReleaseOpenAccountRs 释放OpenAccountRs
func ReleaseOpenAccountRs(v *OpenAccountRs) {
	v.EndDate = ""
	v.StartDate = ""
	v.Url = ""
	poolOpenAccountRs.Put(v)
}
