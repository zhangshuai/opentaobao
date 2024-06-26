package alitripmerchant

import (
	"sync"
)

// DerbyVoucherUniversalDto 结构体
type DerbyVoucherUniversalDto struct {
	// 类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 房型srId
	RoomTypeCode string `json:"room_type_code,omitempty" xml:"room_type_code,omitempty"`
	// 房型rpCode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 权益卡卡号
	MemberVoucherCardId string `json:"member_voucher_card_id,omitempty" xml:"member_voucher_card_id,omitempty"`
	// 权益券id
	MemberVoucherId string `json:"member_voucher_id,omitempty" xml:"member_voucher_id,omitempty"`
	// 折扣百分比
	DiscountOff int64 `json:"discount_off,omitempty" xml:"discount_off,omitempty"`
	// 是否为权益商品
	IsDerbyVoucherRoom bool `json:"is_derby_voucher_room,omitempty" xml:"is_derby_voucher_room,omitempty"`
}

var poolDerbyVoucherUniversalDto = sync.Pool{
	New: func() any {
		return new(DerbyVoucherUniversalDto)
	},
}

// GetDerbyVoucherUniversalDto() 从对象池中获取DerbyVoucherUniversalDto
func GetDerbyVoucherUniversalDto() *DerbyVoucherUniversalDto {
	return poolDerbyVoucherUniversalDto.Get().(*DerbyVoucherUniversalDto)
}

// ReleaseDerbyVoucherUniversalDto 释放DerbyVoucherUniversalDto
func ReleaseDerbyVoucherUniversalDto(v *DerbyVoucherUniversalDto) {
	v.Category = ""
	v.RoomTypeCode = ""
	v.RatePlanCode = ""
	v.MemberVoucherCardId = ""
	v.MemberVoucherId = ""
	v.DiscountOff = 0
	v.IsDerbyVoucherRoom = false
	poolDerbyVoucherUniversalDto.Put(v)
}
