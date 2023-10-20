package simba

import (
	"sync"
)

// InsightWordsAreaDistributeDataDto 结构体
type InsightWordsAreaDistributeDataDto struct {
	// 省名称
	Provincename string `json:"provincename,omitempty" xml:"provincename,omitempty"`
	// 点击转化率
	Coverage string `json:"coverage,omitempty" xml:"coverage,omitempty"`
	// 平均点击花费
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 投入产出比
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 市名称
	Cityname string `json:"cityname,omitempty" xml:"cityname,omitempty"`
	// 关键词
	Bidword string `json:"bidword,omitempty" xml:"bidword,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 展现量
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 间接成交金额
	Indirecttransaction int64 `json:"indirecttransaction,omitempty" xml:"indirecttransaction,omitempty"`
	// 点击量
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 花费，单位（分）
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 直接成交金额
	Directtransaction int64 `json:"directtransaction,omitempty" xml:"directtransaction,omitempty"`
	// 宝贝搜藏数
	Favitemtotal int64 `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
	// 总的成交笔数
	Transactionshippingtotal int64 `json:"transactionshippingtotal,omitempty" xml:"transactionshippingtotal,omitempty"`
	// 总的收藏数，包括宝贝收藏和店铺收藏
	Favtotal int64 `json:"favtotal,omitempty" xml:"favtotal,omitempty"`
	// 成交总金额
	Transactiontotal int64 `json:"transactiontotal,omitempty" xml:"transactiontotal,omitempty"`
	// 间接成交笔数
	Indirecttransactionshipping int64 `json:"indirecttransactionshipping,omitempty" xml:"indirecttransactionshipping,omitempty"`
	// 直接成交笔数
	Directtransactionshipping int64 `json:"directtransactionshipping,omitempty" xml:"directtransactionshipping,omitempty"`
	// 店铺搜藏数
	Favshoptotal int64 `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
	// 竞争度
	Competition int64 `json:"competition,omitempty" xml:"competition,omitempty"`
}

var poolInsightWordsAreaDistributeDataDto = sync.Pool{
	New: func() any {
		return new(InsightWordsAreaDistributeDataDto)
	},
}

// GetInsightWordsAreaDistributeDataDto() 从对象池中获取InsightWordsAreaDistributeDataDto
func GetInsightWordsAreaDistributeDataDto() *InsightWordsAreaDistributeDataDto {
	return poolInsightWordsAreaDistributeDataDto.Get().(*InsightWordsAreaDistributeDataDto)
}

// ReleaseInsightWordsAreaDistributeDataDto 释放InsightWordsAreaDistributeDataDto
func ReleaseInsightWordsAreaDistributeDataDto(v *InsightWordsAreaDistributeDataDto) {
	v.Provincename = ""
	v.Coverage = ""
	v.Cpc = ""
	v.Roi = ""
	v.Cityname = ""
	v.Bidword = ""
	v.Ctr = ""
	v.Impression = 0
	v.Indirecttransaction = 0
	v.Click = 0
	v.Cost = 0
	v.Directtransaction = 0
	v.Favitemtotal = 0
	v.Transactionshippingtotal = 0
	v.Favtotal = 0
	v.Transactiontotal = 0
	v.Indirecttransactionshipping = 0
	v.Directtransactionshipping = 0
	v.Favshoptotal = 0
	v.Competition = 0
	poolInsightWordsAreaDistributeDataDto.Put(v)
}
