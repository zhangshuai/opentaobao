package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费 API请求
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口
*/
type AlibabaAlscCrmRechargeDedutUpdateRequest struct {
    model.Params
    // 入参
    paramDedutOpenReq   *DedutOpenReq
}

// 初始化AlibabaAlscCrmRechargeDedutUpdateRequest对象
func NewAlibabaAlscCrmRechargeDedutUpdateRequest() *AlibabaAlscCrmRechargeDedutUpdateRequest{
    return &AlibabaAlscCrmRechargeDedutUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.dedut.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeDedutUpdateRequest) SetParamDedutOpenReq(paramDedutOpenReq *DedutOpenReq) error {
    r.paramDedutOpenReq = paramDedutOpenReq
    r.Set("param_dedut_open_req", paramDedutOpenReq)
    return nil
}

// ParamDedutOpenReq Getter
func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetParamDedutOpenReq() *DedutOpenReq {
    return r.paramDedutOpenReq
}
