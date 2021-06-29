package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分扣减接口 API请求
taobao.jst.interactive.point.decrease

扣减用户的互动积分
*/
type TaobaoJstInteractivePointDecreaseRequest struct {
    model.Params
    // 扣减的积分值
    amount   int64
    // 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
    operateCode   string
}

// 初始化TaobaoJstInteractivePointDecreaseRequest对象
func NewTaobaoJstInteractivePointDecreaseRequest() *TaobaoJstInteractivePointDecreaseRequest{
    return &TaobaoJstInteractivePointDecreaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointDecreaseRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.point.decrease"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointDecreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Amount Setter
// 扣减的积分值
func (r *TaobaoJstInteractivePointDecreaseRequest) SetAmount(amount int64) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

// Amount Getter
func (r TaobaoJstInteractivePointDecreaseRequest) GetAmount() int64 {
    return r.amount
}
// OperateCode Setter
// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
func (r *TaobaoJstInteractivePointDecreaseRequest) SetOperateCode(operateCode string) error {
    r.operateCode = operateCode
    r.Set("operate_code", operateCode)
    return nil
}

// OperateCode Getter
func (r TaobaoJstInteractivePointDecreaseRequest) GetOperateCode() string {
    return r.operateCode
}