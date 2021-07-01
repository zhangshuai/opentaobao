package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新康众审批门店分账 API请求
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
type TmallServicecenterSettlementStoretransferAuditAPIRequest struct {
    model.Params
    // 审批通过
    _auditPass   bool
    // 工单id
    _workcardId   int64
}

// 初始化TmallServicecenterSettlementStoretransferAuditAPIRequest对象
func NewTmallServicecenterSettlementStoretransferAuditRequest() *TmallServicecenterSettlementStoretransferAuditAPIRequest{
    return &TmallServicecenterSettlementStoretransferAuditAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.settlement.storetransfer.audit"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuditPass Setter
// 审批通过
func (r *TmallServicecenterSettlementStoretransferAuditAPIRequest) SetAuditPass(_auditPass bool) error {
    r._auditPass = _auditPass
    r.Set("audit_pass", _auditPass)
    return nil
}

// AuditPass Getter
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetAuditPass() bool {
    return r._auditPass
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterSettlementStoretransferAuditAPIRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetWorkcardId() int64 {
    return r._workcardId
}