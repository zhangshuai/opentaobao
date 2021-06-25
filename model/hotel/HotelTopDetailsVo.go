package hotel

// HotelTopDetailsVo 
type HotelTopDetailsVo struct {

    // 入住日期
    CheckIn   string `json:"check_in,omitempty"`

    // 离店日期
    CheckOut   string `json:"check_out,omitempty"`

    // 信息
    InfoMsg   string `json:"info_msg,omitempty"`

    // 是否港澳台
    IsGangAoTai   int64 `json:"is_gang_ao_tai,omitempty"`

    // 是否国际
    IsInternational   int64 `json:"is_international,omitempty"`

    // 标签信息
    Labels   []HotelLabelDto `json:"labels,omitempty"`

    // 套餐商品
    PackRoomTypes   []PackageRate `json:"pack_room_types,omitempty"`

    // 房型信息
    RoomTypes   []RoomTypeInfo `json:"room_types,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // price
    Price   int64 `json:"price,omitempty"`

    // 是否显示 附加产品tab页 1-展示 0-否
    ShowHotelPackageTab   int64 `json:"show_hotel_package_tab,omitempty"`

    // galleries
    Galleries   []Gallery `json:"galleries,omitempty"`

    // 立减
    SubtractPrice   string `json:"subtract_price,omitempty"`

    // dynamicUvDesc
    DynamicUvDesc   string `json:"dynamic_uv_desc,omitempty"`

    // rateNumber
    RateNumber   int64 `json:"rate_number,omitempty"`

    // rateScore
    RateScore   string `json:"rate_score,omitempty"`

    // rateScoreDesc
    RateScoreDesc   string `json:"rate_score_desc,omitempty"`

    // rateScoreText
    RateScoreText   string `json:"rate_score_text,omitempty"`

    // scoreText
    ScoreText   string `json:"score_text,omitempty"`

    // star
    Star   string `json:"star,omitempty"`

    // searchId
    SearchId   string `json:"search_id,omitempty"`

}