package alitripmerchant

// CouponParam 结构体
type CouponParam struct {
	// 优惠券模板id集合
	CouponTemplateIdList []int64 `json:"coupon_template_id_list,omitempty" xml:"coupon_template_id_list>int64,omitempty"`
	// 代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 区分版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 优惠价的试单相关参数透传
	DiscountParam *ValidateOrderParam `json:"discount_param,omitempty" xml:"discount_param,omitempty"`
	// 原价的试单相关参数透传
	NoDiscountParam *ValidateOrderParam `json:"no_discount_param,omitempty" xml:"no_discount_param,omitempty"`
	// 权益券相关参数
	DerbyVoucherUniversalDto *DerbyVoucherUniversalDto `json:"derby_voucher_universal_dto,omitempty" xml:"derby_voucher_universal_dto,omitempty"`
}
