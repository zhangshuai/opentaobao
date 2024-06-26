package alsc

import (
	"sync"
)

// DuductedMoneyOpenInfo 结构体
type DuductedMoneyOpenInfo struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 消耗的积分
	ConsumePoint int64 `json:"consume_point,omitempty" xml:"consume_point,omitempty"`
	// 抵扣的金额,单位为分
	DeductedMoney int64 `json:"deducted_money,omitempty" xml:"deducted_money,omitempty"`
	// 是否删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolDuductedMoneyOpenInfo = sync.Pool{
	New: func() any {
		return new(DuductedMoneyOpenInfo)
	},
}

// GetDuductedMoneyOpenInfo() 从对象池中获取DuductedMoneyOpenInfo
func GetDuductedMoneyOpenInfo() *DuductedMoneyOpenInfo {
	return poolDuductedMoneyOpenInfo.Get().(*DuductedMoneyOpenInfo)
}

// ReleaseDuductedMoneyOpenInfo 释放DuductedMoneyOpenInfo
func ReleaseDuductedMoneyOpenInfo(v *DuductedMoneyOpenInfo) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.ConsumePoint = 0
	v.DeductedMoney = 0
	v.Deleted = false
	poolDuductedMoneyOpenInfo.Put(v)
}
