package promotion

import (
	"sync"
)

// CouponResult 结构体
type CouponResult struct {
	// 张三
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// ouida
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// asasa
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 已经发放优惠券的编号
	CouponNumber int64 `json:"coupon_number,omitempty" xml:"coupon_number,omitempty"`
}

var poolCouponResult = sync.Pool{
	New: func() any {
		return new(CouponResult)
	},
}

// GetCouponResult() 从对象池中获取CouponResult
func GetCouponResult() *CouponResult {
	return poolCouponResult.Get().(*CouponResult)
}

// ReleaseCouponResult 释放CouponResult
func ReleaseCouponResult(v *CouponResult) {
	v.BuyerNick = ""
	v.Ouid = ""
	v.OpenUid = ""
	v.CouponNumber = 0
	poolCouponResult.Put(v)
}
