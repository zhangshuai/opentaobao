package cainiaoncwl

// JhItemInfo 
type JhItemInfo struct {
    // 商品名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 外部编号，即商家商品编号
    OuterItemId   string `json:"outer_item_id,omitempty" xml:"outer_item_id,omitempty"`
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 淘系商品号
    TbItemId   int64 `json:"tb_item_id,omitempty" xml:"tb_item_id,omitempty"`
    // 淘系skuid
    TbSkuId   int64 `json:"tb_sku_id,omitempty" xml:"tb_sku_id,omitempty"`
    // 体积，单位，立方分米
    Volume   int64 `json:"volume,omitempty" xml:"volume,omitempty"`
    // 重量，单位，克
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    // 预留字段，活动报名时候配置的，根据包装与包含了几个售卖单位；商家内部有维护，以公司内部为准；
    PacakgeRule   string `json:"pacakge_rule,omitempty" xml:"pacakge_rule,omitempty"`
    // 预留字段，售卖1件相当于几个东西；如果销售商品售卖规则与商家内部规则一致，可以不使用这个字段；
    SellRule   string `json:"sell_rule,omitempty" xml:"sell_rule,omitempty"`
}