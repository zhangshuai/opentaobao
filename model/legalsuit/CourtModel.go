package legalsuit

import (
	"sync"
)

// CourtModel 结构体
type CourtModel struct {
	// 审判员联系方式
	JudgeContact string `json:"judge_contact,omitempty" xml:"judge_contact,omitempty"`
	// 审判员名称
	JudgeName string `json:"judge_name,omitempty" xml:"judge_name,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区/县
	CountyCode string `json:"county_code,omitempty" xml:"county_code,omitempty"`
	// 市
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 受理机关的编号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 受理机关ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCourtModel = sync.Pool{
	New: func() any {
		return new(CourtModel)
	},
}

// GetCourtModel() 从对象池中获取CourtModel
func GetCourtModel() *CourtModel {
	return poolCourtModel.Get().(*CourtModel)
}

// ReleaseCourtModel 释放CourtModel
func ReleaseCourtModel(v *CourtModel) {
	v.JudgeContact = ""
	v.JudgeName = ""
	v.Address = ""
	v.CountyCode = ""
	v.CityCode = ""
	v.ProvinceCode = ""
	v.Name = ""
	v.Code = ""
	v.Id = 0
	poolCourtModel.Put(v)
}
