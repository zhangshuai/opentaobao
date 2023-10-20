package wdk

import (
	"sync"
)

// LimitDto 结构体
type LimitDto struct {
	// 活动总限购
	TotalLimitCnt int64 `json:"total_limit_cnt,omitempty" xml:"total_limit_cnt,omitempty"`
	// 每日总限购
	DailyTotalLimitCnt int64 `json:"daily_total_limit_cnt,omitempty" xml:"daily_total_limit_cnt,omitempty"`
	// 用户总限购
	UserTotalLimitCnt int64 `json:"user_total_limit_cnt,omitempty" xml:"user_total_limit_cnt,omitempty"`
	// 用户每日限购
	UserDailyLimitCnt int64 `json:"user_daily_limit_cnt,omitempty" xml:"user_daily_limit_cnt,omitempty"`
	// 每单限购
	OrderLimitCnt int64 `json:"order_limit_cnt,omitempty" xml:"order_limit_cnt,omitempty"`
}

var poolLimitDto = sync.Pool{
	New: func() any {
		return new(LimitDto)
	},
}

// GetLimitDto() 从对象池中获取LimitDto
func GetLimitDto() *LimitDto {
	return poolLimitDto.Get().(*LimitDto)
}

// ReleaseLimitDto 释放LimitDto
func ReleaseLimitDto(v *LimitDto) {
	v.TotalLimitCnt = 0
	v.DailyTotalLimitCnt = 0
	v.UserTotalLimitCnt = 0
	v.UserDailyLimitCnt = 0
	v.OrderLimitCnt = 0
	poolLimitDto.Put(v)
}
