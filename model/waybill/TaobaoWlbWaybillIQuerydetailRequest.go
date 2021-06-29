package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查面单号状态v1.0 API请求
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。
*/
type TaobaoWlbWaybillIQuerydetailRequest struct {
    model.Params
    // 面单查询请求
    waybillDetailQueryRequest   *WaybillDetailQueryRequest
}

// 初始化TaobaoWlbWaybillIQuerydetailRequest对象
func NewTaobaoWlbWaybillIQuerydetailRequest() *TaobaoWlbWaybillIQuerydetailRequest{
    return &TaobaoWlbWaybillIQuerydetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIQuerydetailRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.querydetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIQuerydetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillDetailQueryRequest Setter
// 面单查询请求
func (r *TaobaoWlbWaybillIQuerydetailRequest) SetWaybillDetailQueryRequest(waybillDetailQueryRequest *WaybillDetailQueryRequest) error {
    r.waybillDetailQueryRequest = waybillDetailQueryRequest
    r.Set("waybill_detail_query_request", waybillDetailQueryRequest)
    return nil
}

// WaybillDetailQueryRequest Getter
func (r TaobaoWlbWaybillIQuerydetailRequest) GetWaybillDetailQueryRequest() *WaybillDetailQueryRequest {
    return r.waybillDetailQueryRequest
}
