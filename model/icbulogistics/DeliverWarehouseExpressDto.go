package icbulogistics

// DeliverWarehouseExpressDto 
type DeliverWarehouseExpressDto struct {
    // 国内快递公司code
    LogisticsCompany   string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
    // 运单号
    TrackingNumbers   []string `json:"tracking_numbers,omitempty" xml:"tracking_numbers>string,omitempty"`
    // 包裹数量
    PackageQuantity   string `json:"package_quantity,omitempty" xml:"package_quantity,omitempty"`
}