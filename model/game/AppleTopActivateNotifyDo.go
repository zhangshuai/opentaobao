package game

import (
	"sync"
)

// AppleTopActivateNotifyDo 结构体
type AppleTopActivateNotifyDo struct {
	// 电子卡卡号
	CardNo string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// 商品编号
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 商户唯一订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 附加信息，后续可以扩展
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 商户上送UCI
	Uci string `json:"uci,omitempty" xml:"uci,omitempty"`
	// 面值，单位分
	FacePrice string `json:"face_price,omitempty" xml:"face_price,omitempty"`
	// 网关订单号
	GatewayOrderNo string `json:"gateway_order_no,omitempty" xml:"gateway_order_no,omitempty"`
	// 电子卡密码
	CardPass string `json:"card_pass,omitempty" xml:"card_pass,omitempty"`
}

var poolAppleTopActivateNotifyDo = sync.Pool{
	New: func() any {
		return new(AppleTopActivateNotifyDo)
	},
}

// GetAppleTopActivateNotifyDo() 从对象池中获取AppleTopActivateNotifyDo
func GetAppleTopActivateNotifyDo() *AppleTopActivateNotifyDo {
	return poolAppleTopActivateNotifyDo.Get().(*AppleTopActivateNotifyDo)
}

// ReleaseAppleTopActivateNotifyDo 释放AppleTopActivateNotifyDo
func ReleaseAppleTopActivateNotifyDo(v *AppleTopActivateNotifyDo) {
	v.CardNo = ""
	v.GoodsId = ""
	v.OrderNo = ""
	v.Memo = ""
	v.Uci = ""
	v.FacePrice = ""
	v.GatewayOrderNo = ""
	v.CardPass = ""
	poolAppleTopActivateNotifyDo.Put(v)
}
