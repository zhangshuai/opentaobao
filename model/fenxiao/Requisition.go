package fenxiao

import (
	"sync"
)

// Requisition 结构体
type Requisition struct {
	// 分销申请加盟时，给供应商的留言
	DistMessage string `json:"dist_message,omitempty" xml:"dist_message,omitempty"`
	// 申请时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 分销商nick
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 开店时间
	DistOpenDate string `json:"dist_open_date,omitempty" xml:"dist_open_date,omitempty"`
	// 店铺地址
	DistShopAddress string `json:"dist_shop_address,omitempty" xml:"dist_shop_address,omitempty"`
	// 主营类目名称
	DistCategoryName string `json:"dist_category_name,omitempty" xml:"dist_category_name,omitempty"`
	// 好评率
	DistAppraise int64 `json:"dist_appraise,omitempty" xml:"dist_appraise,omitempty"`
	// 店铺星级
	DistLevel int64 `json:"dist_level,omitempty" xml:"dist_level,omitempty"`
	// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 主营类目
	DistCategory int64 `json:"dist_category,omitempty" xml:"dist_category,omitempty"`
	// 分销商id
	DistributorId int64 `json:"distributor_id,omitempty" xml:"distributor_id,omitempty"`
	// 合作申请ID
	RequisitionId int64 `json:"requisition_id,omitempty" xml:"requisition_id,omitempty"`
	// 是否消保(0-不是、1-是)
	DistIsXiaobao int64 `json:"dist_is_xiaobao,omitempty" xml:"dist_is_xiaobao,omitempty"`
}

var poolRequisition = sync.Pool{
	New: func() any {
		return new(Requisition)
	},
}

// GetRequisition() 从对象池中获取Requisition
func GetRequisition() *Requisition {
	return poolRequisition.Get().(*Requisition)
}

// ReleaseRequisition 释放Requisition
func ReleaseRequisition(v *Requisition) {
	v.DistMessage = ""
	v.GmtCreate = ""
	v.DistributorNick = ""
	v.DistOpenDate = ""
	v.DistShopAddress = ""
	v.DistCategoryName = ""
	v.DistAppraise = 0
	v.DistLevel = 0
	v.Status = 0
	v.DistCategory = 0
	v.DistributorId = 0
	v.RequisitionId = 0
	v.DistIsXiaobao = 0
	poolRequisition.Put(v)
}
