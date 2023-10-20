package wdk

import (
	"sync"
)

// ActionDto 结构体
type ActionDto struct {
	// 满减金额，单位分
	DecreaseMoney int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
	// 折扣率
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 一口价金额，单位分
	FixPriceMoney int64 `json:"fix_price_money,omitempty" xml:"fix_price_money,omitempty"`
	// 一口价类型 0: 普通一口价 1: 每件或者第n件一口价 2: 整体一口价
	FixPriceType int64 `json:"fix_price_type,omitempty" xml:"fix_price_type,omitempty"`
	// 是否减钱
	Decrease bool `json:"decrease,omitempty" xml:"decrease,omitempty"`
	// 是否打折
	Discount bool `json:"discount,omitempty" xml:"discount,omitempty"`
	// 是否一口价
	FixPrice bool `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
}

var poolActionDto = sync.Pool{
	New: func() any {
		return new(ActionDto)
	},
}

// GetActionDto() 从对象池中获取ActionDto
func GetActionDto() *ActionDto {
	return poolActionDto.Get().(*ActionDto)
}

// ReleaseActionDto 释放ActionDto
func ReleaseActionDto(v *ActionDto) {
	v.DecreaseMoney = 0
	v.DiscountRate = 0
	v.FixPriceMoney = 0
	v.FixPriceType = 0
	v.Decrease = false
	v.Discount = false
	v.FixPrice = false
	poolActionDto.Put(v)
}
