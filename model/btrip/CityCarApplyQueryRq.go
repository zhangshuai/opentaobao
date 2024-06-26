package btrip

import (
	"sync"
)

// CityCarApplyQueryRq 结构体
type CityCarApplyQueryRq struct {
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 审批单创建时间小于值
	CreatedEndAt string `json:"created_end_at,omitempty" xml:"created_end_at,omitempty"`
	// 审批单创建时间大于等于值
	CreatedStartAt string `json:"created_start_at,omitempty" xml:"created_start_at,omitempty"`
	// 三方审批单ID
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 三方员工ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 页码，要求大于等于1，默认1
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 每页数据量，要求大于等于1，默认20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCityCarApplyQueryRq = sync.Pool{
	New: func() any {
		return new(CityCarApplyQueryRq)
	},
}

// GetCityCarApplyQueryRq() 从对象池中获取CityCarApplyQueryRq
func GetCityCarApplyQueryRq() *CityCarApplyQueryRq {
	return poolCityCarApplyQueryRq.Get().(*CityCarApplyQueryRq)
}

// ReleaseCityCarApplyQueryRq 释放CityCarApplyQueryRq
func ReleaseCityCarApplyQueryRq(v *CityCarApplyQueryRq) {
	v.CorpId = ""
	v.CreatedEndAt = ""
	v.CreatedStartAt = ""
	v.ThirdPartApplyId = ""
	v.UserId = ""
	v.PageNumber = 0
	v.PageSize = 0
	poolCityCarApplyQueryRq.Put(v)
}
