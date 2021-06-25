package promotion

// AwardTypeDto 
type AwardTypeDto struct {

    // 奖品创建url
    AwardCreateUrl   string `json:"award_create_url,omitempty"`

    // 奖品类型名称
    AwardTypeName   string `json:"award_type_name,omitempty"`

    // 奖品类型
    AwardType   int64 `json:"award_type,omitempty"`

}