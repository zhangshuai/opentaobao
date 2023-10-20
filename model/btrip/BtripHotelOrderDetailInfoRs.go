package btrip

import (
	"sync"
)

// BtripHotelOrderDetailInfoRs 结构体
type BtripHotelOrderDetailInfoRs struct {
	// 每日房价信息
	DailyPriceInfoList []BtripHotelDailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>btrip_hotel_daily_price_info_dto,omitempty"`
	// 订单酒店信息
	BtripHotelInfo *BtripHotelInfoDto `json:"btrip_hotel_info,omitempty" xml:"btrip_hotel_info,omitempty"`
	// 订单主体信息
	BtripHotelOrderMainInfo *BtripHotelOrderMainInfoDto `json:"btrip_hotel_order_main_info,omitempty" xml:"btrip_hotel_order_main_info,omitempty"`
	// 房型信息
	BtripHotelRoomInfo *BtripHotelRoomInfoDto `json:"btrip_hotel_room_info,omitempty" xml:"btrip_hotel_room_info,omitempty"`
}

var poolBtripHotelOrderDetailInfoRs = sync.Pool{
	New: func() any {
		return new(BtripHotelOrderDetailInfoRs)
	},
}

// GetBtripHotelOrderDetailInfoRs() 从对象池中获取BtripHotelOrderDetailInfoRs
func GetBtripHotelOrderDetailInfoRs() *BtripHotelOrderDetailInfoRs {
	return poolBtripHotelOrderDetailInfoRs.Get().(*BtripHotelOrderDetailInfoRs)
}

// ReleaseBtripHotelOrderDetailInfoRs 释放BtripHotelOrderDetailInfoRs
func ReleaseBtripHotelOrderDetailInfoRs(v *BtripHotelOrderDetailInfoRs) {
	v.DailyPriceInfoList = v.DailyPriceInfoList[:0]
	v.BtripHotelInfo = nil
	v.BtripHotelOrderMainInfo = nil
	v.BtripHotelRoomInfo = nil
	poolBtripHotelOrderDetailInfoRs.Put(v)
}
