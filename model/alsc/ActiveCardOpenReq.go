package alsc

// ActiveCardOpenReq 
type ActiveCardOpenReq struct {

    // 品牌ID,,brand_id与out_brand_id不可同时为空
    BrandId   string `json:"brand_id,omitempty"`

    // 电子卡号
    CardId   string `json:"card_id,omitempty"`

    // 请求ID，用于幂等处理
    RequestId   string `json:"request_id,omitempty"`

    // 店铺ID，shop_id和out_shop_id不可同时为空
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID，shop_id和out_shop_id不可同时为空
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌ID,brand_id与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty"`

}