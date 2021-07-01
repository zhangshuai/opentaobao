package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住订单取消 API请求
taobao.xhotel.order.alipayface.cancel

提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
*/
type TaobaoXhotelOrderAlipayfaceCancelAPIRequest struct {
    model.Params
    // 淘宝订单ID，必选
    _tid   int64
    // 原因描述
    _reasonText   string
    // 外部订单号
    _outId   string
    // 预留后续用
    _notifyUrl   string
    // 请求流水号
    _outUuid   string
}

// 初始化TaobaoXhotelOrderAlipayfaceCancelAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceCancelRequest() *TaobaoXhotelOrderAlipayfaceCancelAPIRequest{
    return &TaobaoXhotelOrderAlipayfaceCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单ID，必选
func (r *TaobaoXhotelOrderAlipayfaceCancelAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetTid() int64 {
    return r._tid
}
// ReasonText Setter
// 原因描述
func (r *TaobaoXhotelOrderAlipayfaceCancelAPIRequest) SetReasonText(_reasonText string) error {
    r._reasonText = _reasonText
    r.Set("reason_text", _reasonText)
    return nil
}

// ReasonText Getter
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetReasonText() string {
    return r._reasonText
}
// OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderAlipayfaceCancelAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetOutId() string {
    return r._outId
}
// NotifyUrl Setter
// 预留后续用
func (r *TaobaoXhotelOrderAlipayfaceCancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
    r._notifyUrl = _notifyUrl
    r.Set("notify_url", _notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetNotifyUrl() string {
    return r._notifyUrl
}
// OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderAlipayfaceCancelAPIRequest) SetOutUuid(_outUuid string) error {
    r._outUuid = _outUuid
    r.Set("out_uuid", _outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelAPIRequest) GetOutUuid() string {
    return r._outUuid
}