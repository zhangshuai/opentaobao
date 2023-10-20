package tbk

// RightsSendRelationRptDto 结构体
type RightsSendRelationRptDto struct {
	// 日期
	Bizdate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 渠道关系id关联的pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 渠道关系id
	Relationid int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 红包发放数量
	Fundnum int64 `json:"fund_num,omitempty" xml:"fund_num,omitempty"`
	// 红包使用次数
	Usenum int64 `json:"use_num,omitempty" xml:"use_num,omitempty"`
}
