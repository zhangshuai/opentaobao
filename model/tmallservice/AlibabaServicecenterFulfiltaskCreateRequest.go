package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合单生成核销单 API请求
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
type AlibabaServicecenterFulfiltaskCreateRequest struct {
    model.Params
    // 工单id列表
    workcardIds   []int64
    // 外部单号
    outerId   string
}

// 初始化AlibabaServicecenterFulfiltaskCreateRequest对象
func NewAlibabaServicecenterFulfiltaskCreateRequest() *AlibabaServicecenterFulfiltaskCreateRequest{
    return &AlibabaServicecenterFulfiltaskCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.fulfiltask.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardIds Setter
// 工单id列表
func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetWorkcardIds(workcardIds []int64) error {
    r.workcardIds = workcardIds
    r.Set("workcard_ids", workcardIds)
    return nil
}

// WorkcardIds Getter
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetWorkcardIds() []int64 {
    return r.workcardIds
}
// OuterId Setter
// 外部单号
func (r *AlibabaServicecenterFulfiltaskCreateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterFulfiltaskCreateRequest) GetOuterId() string {
    return r.outerId
}
