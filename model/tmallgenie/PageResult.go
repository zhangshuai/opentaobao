package tmallgenie

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 结果集
	ResultList []BotSkillsRelInfo `json:"result_list,omitempty" xml:"result_list>bot_skills_rel_info,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 分页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolPageResult = sync.Pool{
	New: func() any {
		return new(PageResult)
	},
}

// GetPageResult() 从对象池中获取PageResult
func GetPageResult() *PageResult {
	return poolPageResult.Get().(*PageResult)
}

// ReleasePageResult 释放PageResult
func ReleasePageResult(v *PageResult) {
	v.ResultList = v.ResultList[:0]
	v.CurrentPage = 0
	v.PageSize = 0
	v.PageCount = 0
	v.TotalCount = 0
	poolPageResult.Put(v)
}
