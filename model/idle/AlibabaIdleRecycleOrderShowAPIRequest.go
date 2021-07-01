package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1.1 API请求
alibaba.idle.recycle.order.show

查询回收订单
*/
type AlibabaIdleRecycleOrderShowAPIRequest struct {
    model.Params
    // 订单号
    _bizOrderId   int64
}

// 初始化AlibabaIdleRecycleOrderShowAPIRequest对象
func NewAlibabaIdleRecycleOrderShowRequest() *AlibabaIdleRecycleOrderShowAPIRequest{
    return &AlibabaIdleRecycleOrderShowAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.show"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderShowAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}