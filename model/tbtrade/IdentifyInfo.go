package tbtrade

// IdentifyInfo 结构体
type IdentifyInfo struct {
	// 三方鉴定物流相关信息
	IdentifyLogisticsInfos []IdentifyLogisticsInfo `json:"identify_logistics_infos,omitempty" xml:"identify_logistics_infos>identify_logistics_info,omitempty"`
	// 三方鉴定服务相关信息
	IdentifyServiceInfos []IdentifyServiceInfo `json:"identify_service_infos,omitempty" xml:"identify_service_infos>identify_service_info,omitempty"`
}
