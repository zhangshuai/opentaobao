package feedflow

// TaobaoFeedflowItemCampaignPageResultDto 
type TaobaoFeedflowItemCampaignPageResultDto struct {
    // 信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 计划列表
    Results   []CampaignDto `json:"results,omitempty" xml:"results>campaign_dto,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 符合条件的计划数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}