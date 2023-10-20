package campus

// ControllerDto 结构体
type ControllerDto struct {
	// guardConfigList
	ConfigList []GuardConfigDto `json:"config_list,omitempty" xml:"config_list>guard_config_dto,omitempty"`
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
	// 控制器名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 控制器设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}