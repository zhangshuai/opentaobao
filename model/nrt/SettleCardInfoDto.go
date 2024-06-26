package nrt

import (
	"sync"
)

// SettleCardInfoDto 结构体
type SettleCardInfoDto struct {
	// 开户支行名
	AccountBranchName string `json:"account_branch_name,omitempty" xml:"account_branch_name,omitempty"`
	// 银行卡号
	AccountNo string `json:"account_no,omitempty" xml:"account_no,omitempty"`
	// 卡类型
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 银行名称
	AccountInstName string `json:"account_inst_name,omitempty" xml:"account_inst_name,omitempty"`
	// 开户行所在地-市
	AccountInstCity string `json:"account_inst_city,omitempty" xml:"account_inst_city,omitempty"`
	// 开户行所在地-省
	AccountInstProvince string `json:"account_inst_province,omitempty" xml:"account_inst_province,omitempty"`
	// 开户行简称缩写
	AccountInstId string `json:"account_inst_id,omitempty" xml:"account_inst_id,omitempty"`
	// 卡户名
	AccountHolderName string `json:"account_holder_name,omitempty" xml:"account_holder_name,omitempty"`
	// 账号使用类型
	UsageType string `json:"usage_type,omitempty" xml:"usage_type,omitempty"`
}

var poolSettleCardInfoDto = sync.Pool{
	New: func() any {
		return new(SettleCardInfoDto)
	},
}

// GetSettleCardInfoDto() 从对象池中获取SettleCardInfoDto
func GetSettleCardInfoDto() *SettleCardInfoDto {
	return poolSettleCardInfoDto.Get().(*SettleCardInfoDto)
}

// ReleaseSettleCardInfoDto 释放SettleCardInfoDto
func ReleaseSettleCardInfoDto(v *SettleCardInfoDto) {
	v.AccountBranchName = ""
	v.AccountNo = ""
	v.AccountType = ""
	v.AccountInstName = ""
	v.AccountInstCity = ""
	v.AccountInstProvince = ""
	v.AccountInstId = ""
	v.AccountHolderName = ""
	v.UsageType = ""
	poolSettleCardInfoDto.Put(v)
}
