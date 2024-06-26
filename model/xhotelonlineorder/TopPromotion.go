package xhotelonlineorder

import (
	"sync"
)

// TopPromotion 结构体
type TopPromotion struct {
	// topCreateOrderPromotion
	TopCreateOrderPromotion *TopOrderPromotionDetail `json:"top_create_order_promotion,omitempty" xml:"top_create_order_promotion,omitempty"`
	// topCurrentOrderPromotion
	TopCurrentOrderPromotion *TopOrderPromotionDetail `json:"top_current_order_promotion,omitempty" xml:"top_current_order_promotion,omitempty"`
	// topOrderPromotionExtend
	TopOrderPromotionExtend *TopOrderPromotionExtend `json:"top_order_promotion_extend,omitempty" xml:"top_order_promotion_extend,omitempty"`
}

var poolTopPromotion = sync.Pool{
	New: func() any {
		return new(TopPromotion)
	},
}

// GetTopPromotion() 从对象池中获取TopPromotion
func GetTopPromotion() *TopPromotion {
	return poolTopPromotion.Get().(*TopPromotion)
}

// ReleaseTopPromotion 释放TopPromotion
func ReleaseTopPromotion(v *TopPromotion) {
	v.TopCreateOrderPromotion = nil
	v.TopCurrentOrderPromotion = nil
	v.TopOrderPromotionExtend = nil
	poolTopPromotion.Put(v)
}
