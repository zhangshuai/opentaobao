package xhotelonlineorder

import (
	"sync"
)

// DailyPriceTo 结构体
type DailyPriceTo struct {
	// 币种名称
	CurrencyCodeName string `json:"currency_code_name,omitempty" xml:"currency_code_name,omitempty"`
	// 日期
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 外币价格
	ForeignCurrencyPrice int64 `json:"foreign_currency_price,omitempty" xml:"foreign_currency_price,omitempty"`
	// 积分数
	Point int64 `json:"point,omitempty" xml:"point,omitempty"`
	// 淘里程数
	Mileage int64 `json:"mileage,omitempty" xml:"mileage,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 税费
	TaxPrice int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// 早餐数
	BreakFastCount int64 `json:"break_fast_count,omitempty" xml:"break_fast_count,omitempty"`
	// 优惠后金额
	AfterPromotionPrice int64 `json:"after_promotion_price,omitempty" xml:"after_promotion_price,omitempty"`
	// 基础金额
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 税费外币金额
	TaxForeignPrice int64 `json:"tax_foreign_price,omitempty" xml:"tax_foreign_price,omitempty"`
	// servicePrice
	ServicePrice int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// serviceForeignPrice
	ServiceForeignPrice int64 `json:"service_foreign_price,omitempty" xml:"service_foreign_price,omitempty"`
	// 外币底价
	BaseForeignPrice int64 `json:"base_foreign_price,omitempty" xml:"base_foreign_price,omitempty"`
}

var poolDailyPriceTo = sync.Pool{
	New: func() any {
		return new(DailyPriceTo)
	},
}

// GetDailyPriceTo() 从对象池中获取DailyPriceTo
func GetDailyPriceTo() *DailyPriceTo {
	return poolDailyPriceTo.Get().(*DailyPriceTo)
}

// ReleaseDailyPriceTo 释放DailyPriceTo
func ReleaseDailyPriceTo(v *DailyPriceTo) {
	v.CurrencyCodeName = ""
	v.Day = ""
	v.ForeignCurrencyPrice = 0
	v.Point = 0
	v.Mileage = 0
	v.Price = 0
	v.TaxPrice = 0
	v.BreakFastCount = 0
	v.AfterPromotionPrice = 0
	v.BasePrice = 0
	v.TaxForeignPrice = 0
	v.ServicePrice = 0
	v.ServiceForeignPrice = 0
	v.BaseForeignPrice = 0
	poolDailyPriceTo.Put(v)
}
