package wdk

// SeasonVersionParam 
type SeasonVersionParam struct {

    // 商家档期号
    OutSeasonId   string `json:"out_season_id,omitempty"`

    // 参与换挡的门店列表
    ShopIds   []String `json:"shop_ids,omitempty"`

}