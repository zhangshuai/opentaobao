package bus

import (
	"sync"
)

// BusCityChangeDto 结构体
type BusCityChangeDto struct {
	// 城市拼音
	CityFullPinyin string `json:"city_full_pinyin,omitempty" xml:"city_full_pinyin,omitempty"`
	// 城市名
	StartCityName string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// 是否开通标示，0：已开通 1：未开通
	TypeNo int64 `json:"type_no,omitempty" xml:"type_no,omitempty"`
	// 是否可售标示，0：暂停售票 1：可售
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolBusCityChangeDto = sync.Pool{
	New: func() any {
		return new(BusCityChangeDto)
	},
}

// GetBusCityChangeDto() 从对象池中获取BusCityChangeDto
func GetBusCityChangeDto() *BusCityChangeDto {
	return poolBusCityChangeDto.Get().(*BusCityChangeDto)
}

// ReleaseBusCityChangeDto 释放BusCityChangeDto
func ReleaseBusCityChangeDto(v *BusCityChangeDto) {
	v.CityFullPinyin = ""
	v.StartCityName = ""
	v.TypeNo = 0
	v.Status = 0
	poolBusCityChangeDto.Put(v)
}
