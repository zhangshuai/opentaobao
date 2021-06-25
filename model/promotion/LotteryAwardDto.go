package promotion

// LotteryAwardDto 
type LotteryAwardDto struct {

    // totalResCount
    TotalResCount   int64 `json:"total_res_count,omitempty"`

    // awardDetailUrl
    AwardDetailUrl   string `json:"award_detail_url,omitempty"`

    // useEndDate
    UseEndDate   string `json:"use_end_date,omitempty"`

    // useStartDate
    UseStartDate   string `json:"use_start_date,omitempty"`

    // endDate
    EndDate   string `json:"end_date,omitempty"`

    // startDate
    StartDate   string `json:"start_date,omitempty"`

    // 币种
    Currency   string `json:"currency,omitempty"`

    // awardPrice
    AwardPrice   int64 `json:"award_price,omitempty"`

    // startFee
    StartFee   int64 `json:"start_fee,omitempty"`

    // displayName
    DisplayName   string `json:"display_name,omitempty"`

    // name
    Name   string `json:"name,omitempty"`

    // awardTypeName
    AwardTypeName   string `json:"award_type_name,omitempty"`

    // canUseResCount
    CanUseResCount   int64 `json:"can_use_res_count,omitempty"`

    // condition
    Condition   string `json:"condition,omitempty"`

    // 币种符号
    CurrencySign   string `json:"currency_sign,omitempty"`

    // pictUrl
    PictUrl   string `json:"pict_url,omitempty"`

    // awardInstId
    AwardInstId   string `json:"award_inst_id,omitempty"`

    // awardType
    AwardType   int64 `json:"award_type,omitempty"`

    // id
    Id   int64 `json:"id,omitempty"`

    // lotteryActivityId
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`

}