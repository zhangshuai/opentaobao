package flight

// IntentionListRequestDto 结构体
type IntentionListRequestDto struct {
	// 市场航司
	MarketingAirline []string `json:"marketing_airline,omitempty" xml:"marketing_airline>string,omitempty"`
	// 承运航班号
	OperatingFlightNo []string `json:"operating_flight_no,omitempty" xml:"operating_flight_no>string,omitempty"`
	// 起飞机场
	DepAirportCode []string `json:"dep_airport_code,omitempty" xml:"dep_airport_code>string,omitempty"`
	// 市场航班号
	MarketingFlightNo []string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no>string,omitempty"`
	// 抵达机场
	ArrAirportCode []string `json:"arr_airport_code,omitempty" xml:"arr_airport_code>string,omitempty"`
	// 承运航司
	OperatingAirline []string `json:"operating_airline,omitempty" xml:"operating_airline>string,omitempty"`
	// 截止起飞时间
	DepTimeEnd string `json:"dep_time_end,omitempty" xml:"dep_time_end,omitempty"`
	// 开始起飞时间
	DepTimeStart string `json:"dep_time_start,omitempty" xml:"dep_time_start,omitempty"`
	// 页数
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 是否查询可确认意向单
	OnlyInit bool `json:"only_init,omitempty" xml:"only_init,omitempty"`
}
