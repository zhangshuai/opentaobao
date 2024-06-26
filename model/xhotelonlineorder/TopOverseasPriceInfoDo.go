package xhotelonlineorder

import (
	"sync"
)

// TopOverseasPriceInfoDo 结构体
type TopOverseasPriceInfoDo struct {
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 下单时汇率
	BookingRate string `json:"booking_rate,omitempty" xml:"booking_rate,omitempty"`
	// 当前汇率
	Rate string `json:"rate,omitempty" xml:"rate,omitempty"`
	// 离店日汇率
	CheckoutRate string `json:"checkout_rate,omitempty" xml:"checkout_rate,omitempty"`
	// 产品售价,单位(元*100),包含税服务费等
	BookingProductSalePrice int64 `json:"booking_product_sale_price,omitempty" xml:"booking_product_sale_price,omitempty"`
	// 买家应付包含杂费,单位(元*100)
	BuyerPayment int64 `json:"buyer_payment,omitempty" xml:"buyer_payment,omitempty"`
	// 税费,单位(元*100)
	TaxPrice int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// 卖家应收包含杂费,单位(元*100)
	SellerPayment int64 `json:"seller_payment,omitempty" xml:"seller_payment,omitempty"`
	// 结账杂费,单位(元*100)
	OtherFee int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
	// 服务费,单位(元*100)
	BookingServicePrice int64 `json:"booking_service_price,omitempty" xml:"booking_service_price,omitempty"`
	// 下单时刻卖家应收
	BookSellerPayment int64 `json:"book_seller_payment,omitempty" xml:"book_seller_payment,omitempty"`
	// 税费,单位(元*100)
	BookingTaxPrice int64 `json:"booking_tax_price,omitempty" xml:"booking_tax_price,omitempty"`
	// 总价包含杂费,单位(元*100)
	TotalPayment int64 `json:"total_payment,omitempty" xml:"total_payment,omitempty"`
	// 担保金(元*100)
	FundPrice int64 `json:"fund_price,omitempty" xml:"fund_price,omitempty"`
	// 结账产品售价,单位(元*100),包含税服务费等
	SettleProductSalePrice int64 `json:"settle_product_sale_price,omitempty" xml:"settle_product_sale_price,omitempty"`
	// 服务费,单位(元*100)
	ServicePrice int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// 日历价格
	CreateOrderDailyPrice *HbsDailyPriceDo `json:"create_order_daily_price,omitempty" xml:"create_order_daily_price,omitempty"`
	// bookingRoomsPrice
	BookingRoomsPrice int64 `json:"booking_rooms_price,omitempty" xml:"booking_rooms_price,omitempty"`
	// settleOrderDailyPrice
	SettleOrderDailyPrice *HbsDailyPriceDo `json:"settle_order_daily_price,omitempty" xml:"settle_order_daily_price,omitempty"`
	// 房费
	RoomsPrice int64 `json:"rooms_price,omitempty" xml:"rooms_price,omitempty"`
	// 底价房费
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 下单时底价房费
	BookingBasePrice int64 `json:"booking_base_price,omitempty" xml:"booking_base_price,omitempty"`
}

var poolTopOverseasPriceInfoDo = sync.Pool{
	New: func() any {
		return new(TopOverseasPriceInfoDo)
	},
}

// GetTopOverseasPriceInfoDo() 从对象池中获取TopOverseasPriceInfoDo
func GetTopOverseasPriceInfoDo() *TopOverseasPriceInfoDo {
	return poolTopOverseasPriceInfoDo.Get().(*TopOverseasPriceInfoDo)
}

// ReleaseTopOverseasPriceInfoDo 释放TopOverseasPriceInfoDo
func ReleaseTopOverseasPriceInfoDo(v *TopOverseasPriceInfoDo) {
	v.CurrencyCode = ""
	v.BookingRate = ""
	v.Rate = ""
	v.CheckoutRate = ""
	v.BookingProductSalePrice = 0
	v.BuyerPayment = 0
	v.TaxPrice = 0
	v.SellerPayment = 0
	v.OtherFee = 0
	v.BookingServicePrice = 0
	v.BookSellerPayment = 0
	v.BookingTaxPrice = 0
	v.TotalPayment = 0
	v.FundPrice = 0
	v.SettleProductSalePrice = 0
	v.ServicePrice = 0
	v.CreateOrderDailyPrice = nil
	v.BookingRoomsPrice = 0
	v.SettleOrderDailyPrice = nil
	v.RoomsPrice = 0
	v.BasePrice = 0
	v.BookingBasePrice = 0
	poolTopOverseasPriceInfoDo.Put(v)
}
