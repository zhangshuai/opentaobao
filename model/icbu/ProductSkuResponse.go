package icbu

// ProductSkuResponse 
type ProductSkuResponse struct {
    // SKU使用的属性
    SkuAttributes   []SkuAttribute `json:"sku_attributes,omitempty" xml:"sku_attributes>sku_attribute,omitempty"`
    // SKU定义
    Skus   []SkuDefinition `json:"skus,omitempty" xml:"skus>sku_definition,omitempty"`
}
