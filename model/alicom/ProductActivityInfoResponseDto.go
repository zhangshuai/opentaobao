package alicom

// ProductActivityInfoResponseDto 
type ProductActivityInfoResponseDto struct {

    // 产品ID
    ProductId   string `json:"product_id,omitempty"`

    // 产品名称
    ProductName   string `json:"product_name,omitempty"`

    // 卖家名称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 金额
    Price   string `json:"price,omitempty"`

    // 活动ID
    ActivityId   string `json:"activity_id,omitempty"`

    // 活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // activityGiftInfos
    ActivityGiftInfos   *ActivityGiftInfos `json:"activity_gift_infos,omitempty"`

}