package btrip

import (
	"sync"
)

// FlightSearchRq 结构体
type FlightSearchRq struct {
	// 查询的行程列表
	OdInfoList []OdInfoRq `json:"od_info_list,omitempty" xml:"od_info_list>od_info_rq,omitempty"`
	// 乘客类型和数量
	PassengerQuantityList []PassengerQuantityRq `json:"passenger_quantity_list,omitempty" xml:"passenger_quantity_list>passenger_quantity_rq,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 飞机信息搜索
	TripPreference *TripPreferenceRq `json:"trip_preference,omitempty" xml:"trip_preference,omitempty"`
}

var poolFlightSearchRq = sync.Pool{
	New: func() any {
		return new(FlightSearchRq)
	},
}

// GetFlightSearchRq() 从对象池中获取FlightSearchRq
func GetFlightSearchRq() *FlightSearchRq {
	return poolFlightSearchRq.Get().(*FlightSearchRq)
}

// ReleaseFlightSearchRq 释放FlightSearchRq
func ReleaseFlightSearchRq(v *FlightSearchRq) {
	v.OdInfoList = v.OdInfoList[:0]
	v.PassengerQuantityList = v.PassengerQuantityList[:0]
	v.CorpId = ""
	v.TripPreference = nil
	poolFlightSearchRq.Put(v)
}
