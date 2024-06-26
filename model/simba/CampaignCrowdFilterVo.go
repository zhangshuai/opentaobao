package simba

import (
	"sync"
)

// CampaignCrowdFilterVo 结构体
type CampaignCrowdFilterVo struct {
	// 行为,purchase:购买人群,cart:加购人群,collection:收藏人群,arrival:进店人群
	Behavior string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// 时间窗口,7:7天,15:15天,30:30天,90:90天,180:180天,365:365天
	Window string `json:"window,omitempty" xml:"window,omitempty"`
}

var poolCampaignCrowdFilterVo = sync.Pool{
	New: func() any {
		return new(CampaignCrowdFilterVo)
	},
}

// GetCampaignCrowdFilterVo() 从对象池中获取CampaignCrowdFilterVo
func GetCampaignCrowdFilterVo() *CampaignCrowdFilterVo {
	return poolCampaignCrowdFilterVo.Get().(*CampaignCrowdFilterVo)
}

// ReleaseCampaignCrowdFilterVo 释放CampaignCrowdFilterVo
func ReleaseCampaignCrowdFilterVo(v *CampaignCrowdFilterVo) {
	v.Behavior = ""
	v.Window = ""
	poolCampaignCrowdFilterVo.Put(v)
}
