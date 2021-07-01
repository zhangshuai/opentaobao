package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-无需物流发货 API请求
alibaba.lst.logistics.notrace.send

异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsNotraceSendAPIRequest struct {
    model.Params
    // 入参
    _param   *SendDummyOrderParam
}

// 初始化AlibabaLstLogisticsNotraceSendAPIRequest对象
func NewAlibabaLstLogisticsNotraceSendRequest() *AlibabaLstLogisticsNotraceSendAPIRequest{
    return &AlibabaLstLogisticsNotraceSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsNotraceSendAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.notrace.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsNotraceSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstLogisticsNotraceSendAPIRequest) SetParam(_param *SendDummyOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstLogisticsNotraceSendAPIRequest) GetParam() *SendDummyOrderParam {
    return r._param
}