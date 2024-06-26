package icbulogistics

import (
	"sync"
)

// RegionEntity 结构体
type RegionEntity struct {
	// 子节点列表
	Childrens []Division `json:"childrens,omitempty" xml:"childrens>division,omitempty"`
	// 节点名称拼音
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 中文名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 上级节点名称
	ParentName string `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
	// 邮编
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 节点ID
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 上级节点ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 层级
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolRegionEntity = sync.Pool{
	New: func() any {
		return new(RegionEntity)
	},
}

// GetRegionEntity() 从对象池中获取RegionEntity
func GetRegionEntity() *RegionEntity {
	return poolRegionEntity.Get().(*RegionEntity)
}

// ReleaseRegionEntity 释放RegionEntity
func ReleaseRegionEntity(v *RegionEntity) {
	v.Childrens = v.Childrens[:0]
	v.Pinyin = ""
	v.Name = ""
	v.ParentName = ""
	v.Zip = ""
	v.Address = ""
	v.AreaId = 0
	v.ParentId = 0
	v.Level = 0
	v.Id = 0
	poolRegionEntity.Put(v)
}
