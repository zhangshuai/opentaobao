package icburfq

import (
	"sync"
)

// RecommendRfqDto 结构体
type RecommendRfqDto struct {
	// 国家全称
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 国家简称
	CountrySimple string `json:"country_simple,omitempty" xml:"country_simple,omitempty"`
	// 发布时间
	DatePostStr string `json:"date_post_str,omitempty" xml:"date_post_str,omitempty"`
	// RFQ详情
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// rfqID
	RfqId string `json:"rfq_id,omitempty" xml:"rfq_id,omitempty"`
	// RFQ标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 数量单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 发布时间戳
	DatePost int64 `json:"date_post,omitempty" xml:"date_post,omitempty"`
	// 剩余报价数
	LeftCount int64 `json:"left_count,omitempty" xml:"left_count,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 是否有读过
	HasRead bool `json:"has_read,omitempty" xml:"has_read,omitempty"`
	// 是否有图片
	HaveImage bool `json:"have_image,omitempty" xml:"have_image,omitempty"`
}

var poolRecommendRfqDto = sync.Pool{
	New: func() any {
		return new(RecommendRfqDto)
	},
}

// GetRecommendRfqDto() 从对象池中获取RecommendRfqDto
func GetRecommendRfqDto() *RecommendRfqDto {
	return poolRecommendRfqDto.Get().(*RecommendRfqDto)
}

// ReleaseRecommendRfqDto 释放RecommendRfqDto
func ReleaseRecommendRfqDto(v *RecommendRfqDto) {
	v.Country = ""
	v.CountrySimple = ""
	v.DatePostStr = ""
	v.Detail = ""
	v.ImageUrl = ""
	v.RfqId = ""
	v.Subject = ""
	v.QuantityUnit = ""
	v.DatePost = 0
	v.LeftCount = 0
	v.Quantity = 0
	v.HasRead = false
	v.HaveImage = false
	poolRecommendRfqDto.Put(v)
}
