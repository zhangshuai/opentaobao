package xhotelitem

import (
	"sync"
)

// PromoInfo 结构体
type PromoInfo struct {
	// 今夜特惠
	TonightDiscount *TonightDiscount `json:"tonight_discount,omitempty" xml:"tonight_discount,omitempty"`
	// 连住优惠
	LongOrderInfo *LongOrderInfo `json:"long_order_info,omitempty" xml:"long_order_info,omitempty"`
	// 早定优惠
	EarlyBookingInfo *EarlyBookingInfo `json:"early_booking_info,omitempty" xml:"early_booking_info,omitempty"`
	// 天天特惠
	DailyBookingInfo *DailyBookingInfo `json:"daily_booking_info,omitempty" xml:"daily_booking_info,omitempty"`
	// 民宿优惠
	GeneralBookingInfo *GeneralBookingInfo `json:"general_booking_info,omitempty" xml:"general_booking_info,omitempty"`
}

var poolPromoInfo = sync.Pool{
	New: func() any {
		return new(PromoInfo)
	},
}

// GetPromoInfo() 从对象池中获取PromoInfo
func GetPromoInfo() *PromoInfo {
	return poolPromoInfo.Get().(*PromoInfo)
}

// ReleasePromoInfo 释放PromoInfo
func ReleasePromoInfo(v *PromoInfo) {
	v.TonightDiscount = nil
	v.LongOrderInfo = nil
	v.EarlyBookingInfo = nil
	v.DailyBookingInfo = nil
	v.GeneralBookingInfo = nil
	poolPromoInfo.Put(v)
}
