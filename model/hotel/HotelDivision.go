package hotel

import (
	"sync"
)

// HotelDivision 结构体
type HotelDivision struct {
	// 城市名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 0：国内；1：国外
	Region int64 `json:"region,omitempty" xml:"region,omitempty"`
	// 城市编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 层级，1：国家，2：州省，3：城市，4：区县
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 类型，0：普通，1：景点,cityTag，如千岛湖
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 父行政区编码，例如杭州市(330100)的父行政区为浙江省(330000)。注意：如果parent_code的值为0，表示该行政区为最高级别的行政区，没有父行政区。
	ParentCode int64 `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// code是否可用来搜索酒店，true：code可直接用于搜索酒店，false: code不可直接用于搜索酒店。
	Searchable bool `json:"searchable,omitempty" xml:"searchable,omitempty"`
}

var poolHotelDivision = sync.Pool{
	New: func() any {
		return new(HotelDivision)
	},
}

// GetHotelDivision() 从对象池中获取HotelDivision
func GetHotelDivision() *HotelDivision {
	return poolHotelDivision.Get().(*HotelDivision)
}

// ReleaseHotelDivision 释放HotelDivision
func ReleaseHotelDivision(v *HotelDivision) {
	v.Name = ""
	v.Region = 0
	v.Code = 0
	v.Level = 0
	v.Type = 0
	v.ParentCode = 0
	v.Searchable = false
	poolHotelDivision.Put(v)
}
