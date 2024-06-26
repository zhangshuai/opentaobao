package feedflow

import (
	"sync"
)

// ItemQueryDto 结构体
type ItemQueryDto struct {
	// 商品id列表
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 分页页码，不得超过20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolItemQueryDto = sync.Pool{
	New: func() any {
		return new(ItemQueryDto)
	},
}

// GetItemQueryDto() 从对象池中获取ItemQueryDto
func GetItemQueryDto() *ItemQueryDto {
	return poolItemQueryDto.Get().(*ItemQueryDto)
}

// ReleaseItemQueryDto 释放ItemQueryDto
func ReleaseItemQueryDto(v *ItemQueryDto) {
	v.ItemIdList = v.ItemIdList[:0]
	v.Title = ""
	v.CampaignId = 0
	v.PageSize = 0
	v.CurrentPage = 0
	poolItemQueryDto.Put(v)
}
