package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态查询接口 API请求
alibaba.ele.enterprise.ordernew.getstatus

订单状态查询接口
*/
type AlibabaEleEnterpriseOrdernewGetstatusAPIRequest struct {
    model.Params
    // 订单号
    _elemeOrderId   string
}

// 初始化AlibabaEleEnterpriseOrdernewGetstatusAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGetstatusRequest() *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest{
    return &AlibabaEleEnterpriseOrdernewGetstatusAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.getstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ElemeOrderId Setter
// 订单号
func (r *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) SetElemeOrderId(_elemeOrderId string) error {
    r._elemeOrderId = _elemeOrderId
    r.Set("eleme_order_id", _elemeOrderId)
    return nil
}

// ElemeOrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetElemeOrderId() string {
    return r._elemeOrderId
}