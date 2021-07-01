package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CP兜底后指定接单的小件员 API请求
cainiao.guoguo.cp.backup.assigncourier

CP兜底后指定接单的小件员；CP改派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierAPIRequest struct {
    model.Params
    // 小件员所在公司编号
    _cpCode   string
    // 小件员员工编号
    _cpUserId   string
    // LP订单号
    _lpCode   string
    // 任务ID
    _taskId   int64
    // 指派/改派原因编码
    _assignReasonCode   string
    // 指派/改派原因
    _assignReason   string
    // 小件员手机号
    _mobile   string
}

// 初始化CainiaoGuoguoCpBackupAssigncourierAPIRequest对象
func NewCainiaoGuoguoCpBackupAssigncourierRequest() *CainiaoGuoguoCpBackupAssigncourierAPIRequest{
    return &CainiaoGuoguoCpBackupAssigncourierAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.backup.assigncourier"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetCpCode() string {
    return r._cpCode
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetCpUserId(_cpUserId string) error {
    r._cpUserId = _cpUserId
    r.Set("cp_user_id", _cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetCpUserId() string {
    return r._cpUserId
}
// LpCode Setter
// LP订单号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetLpCode(_lpCode string) error {
    r._lpCode = _lpCode
    r.Set("lp_code", _lpCode)
    return nil
}

// LpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetLpCode() string {
    return r._lpCode
}
// TaskId Setter
// 任务ID
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetTaskId(_taskId int64) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetTaskId() int64 {
    return r._taskId
}
// AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetAssignReasonCode(_assignReasonCode string) error {
    r._assignReasonCode = _assignReasonCode
    r.Set("assign_reason_code", _assignReasonCode)
    return nil
}

// AssignReasonCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetAssignReasonCode() string {
    return r._assignReasonCode
}
// AssignReason Setter
// 指派/改派原因
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetAssignReason(_assignReason string) error {
    r._assignReason = _assignReason
    r.Set("assign_reason", _assignReason)
    return nil
}

// AssignReason Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetAssignReason() string {
    return r._assignReason
}
// Mobile Setter
// 小件员手机号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetMobile() string {
    return r._mobile
}