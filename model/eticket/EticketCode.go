package eticket

// EticketCode 
type EticketCode struct {

    // 订单ID
    OrderId   int64 `json:"order_id,omitempty"`

    // 电子凭证码
    Code   string `json:"code,omitempty"`

    // 可用数量
    Num   int64 `json:"num,omitempty"`

    // 码状态
    Status   int64 `json:"status,omitempty"`

    // 二维码的图片地址
    QrcodeUrl   string `json:"qrcode_url,omitempty"`

}