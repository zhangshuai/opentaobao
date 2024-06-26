package btrip

import (
	"sync"
)

// BusSearchRs 结构体
type BusSearchRs struct {
	// 查询结果
	BusLines []BusLineInfoVo `json:"bus_lines,omitempty" xml:"bus_lines>bus_line_info_vo,omitempty"`
	// 所有出发车站
	DepStations []string `json:"dep_stations,omitempty" xml:"dep_stations>string,omitempty"`
	// 推荐的路线
	RecommendRoutes []RouteVo `json:"recommend_routes,omitempty" xml:"recommend_routes>route_vo,omitempty"`
	// 车站信息集合
	StationLatitudeLongitudes []StationLatitudeLongitudeVo `json:"station_latitude_longitudes,omitempty" xml:"station_latitude_longitudes>station_latitude_longitude_vo,omitempty"`
	// 到达站列表
	ToStationNames []string `json:"to_station_names,omitempty" xml:"to_station_names>string,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 是否放大搜索  0 没有 1 放大出发 2 放大到达
	EnLarge int64 `json:"en_large,omitempty" xml:"en_large,omitempty"`
	// 重名的城市
	NameSameCity *NameSameCityVo `json:"name_same_city,omitempty" xml:"name_same_city,omitempty"`
	// 车次数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否是预约购票订单
	PreOrder bool `json:"pre_order,omitempty" xml:"pre_order,omitempty"`
}

var poolBusSearchRs = sync.Pool{
	New: func() any {
		return new(BusSearchRs)
	},
}

// GetBusSearchRs() 从对象池中获取BusSearchRs
func GetBusSearchRs() *BusSearchRs {
	return poolBusSearchRs.Get().(*BusSearchRs)
}

// ReleaseBusSearchRs 释放BusSearchRs
func ReleaseBusSearchRs(v *BusSearchRs) {
	v.BusLines = v.BusLines[:0]
	v.DepStations = v.DepStations[:0]
	v.RecommendRoutes = v.RecommendRoutes[:0]
	v.StationLatitudeLongitudes = v.StationLatitudeLongitudes[:0]
	v.ToStationNames = v.ToStationNames[:0]
	v.ArrCity = ""
	v.DepCity = ""
	v.DepDate = ""
	v.EnLarge = 0
	v.NameSameCity = nil
	v.Total = 0
	v.PreOrder = false
	poolBusSearchRs.Put(v)
}
