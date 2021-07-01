package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的核销接口 API请求
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
*/
type TaobaoVmarketEticketAuthConsumeAPIRequest struct {
    model.Params
    // 核销的码，只支持单个码，多个码核销需要多次调用
    _verifyCode   string
    // 核销份数
    _consumeNum   int64
    // 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
    _operatorid   string
    // 自定义核销流水号，需要小于等于100个字符(a-zA-Z0-9_)
    _serialNum   string
    // 网点ID,网点授权核销时，必须传入；其他核销方式可不传
    _storeid   string
}

// 初始化TaobaoVmarketEticketAuthConsumeAPIRequest对象
func NewTaobaoVmarketEticketAuthConsumeRequest() *TaobaoVmarketEticketAuthConsumeAPIRequest{
    return &TaobaoVmarketEticketAuthConsumeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.auth.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyCode Setter
// 核销的码，只支持单个码，多个码核销需要多次调用
func (r *TaobaoVmarketEticketAuthConsumeAPIRequest) SetVerifyCode(_verifyCode string) error {
    r._verifyCode = _verifyCode
    r.Set("verify_code", _verifyCode)
    return nil
}

// VerifyCode Getter
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetVerifyCode() string {
    return r._verifyCode
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoVmarketEticketAuthConsumeAPIRequest) SetConsumeNum(_consumeNum int64) error {
    r._consumeNum = _consumeNum
    r.Set("consume_num", _consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetConsumeNum() int64 {
    return r._consumeNum
}
// Operatorid Setter
// 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
func (r *TaobaoVmarketEticketAuthConsumeAPIRequest) SetOperatorid(_operatorid string) error {
    r._operatorid = _operatorid
    r.Set("operatorid", _operatorid)
    return nil
}

// Operatorid Getter
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetOperatorid() string {
    return r._operatorid
}
// SerialNum Setter
// 自定义核销流水号，需要小于等于100个字符(a-zA-Z0-9_)
func (r *TaobaoVmarketEticketAuthConsumeAPIRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetSerialNum() string {
    return r._serialNum
}
// Storeid Setter
// 网点ID,网点授权核销时，必须传入；其他核销方式可不传
func (r *TaobaoVmarketEticketAuthConsumeAPIRequest) SetStoreid(_storeid string) error {
    r._storeid = _storeid
    r.Set("storeid", _storeid)
    return nil
}

// Storeid Getter
func (r TaobaoVmarketEticketAuthConsumeAPIRequest) GetStoreid() string {
    return r._storeid
}