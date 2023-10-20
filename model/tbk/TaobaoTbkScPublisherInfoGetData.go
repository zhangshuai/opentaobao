package tbk

// TaobaotbkscpublisherinfogetData 结构体
type TaobaotbkscpublisherinfogetData struct {
	// 共享字段 - 渠道或会员列表
	Inviterlist []TaobaotbkscpublisherinfogetMapData `json:"inviter_list,omitempty" xml:"inviter_list>taobaotbkscpublisherinfoget_map_data,omitempty"`
	// 渠道专属pidList
	Rootpidchannellist []string `json:"root_pid_channel_list,omitempty" xml:"root_pid_channel_list>string,omitempty"`
	// 共享字段 - 总记录数
	Totalcount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
