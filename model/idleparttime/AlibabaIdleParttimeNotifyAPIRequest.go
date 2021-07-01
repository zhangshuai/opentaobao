package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职通知接口 API请求
alibaba.idle.parttime.notify

服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
*/
type AlibabaIdleParttimeNotifyAPIRequest struct {
    model.Params
    // 实时同步类型, 0: 岗位状态, 1: 录取状态
    _type   int64
    // 岗位: 0关闭 ; 录取: 0不录取, 1录取
    _status   int64
    // 岗位id
    _jobId   int64
    // 用户id
    _userId   int64
    // 同步时间
    _syncTime   int64
    // 报名id
    _applyId   int64
    // 通知消息
    _message   string
}

// 初始化AlibabaIdleParttimeNotifyAPIRequest对象
func NewAlibabaIdleParttimeNotifyRequest() *AlibabaIdleParttimeNotifyAPIRequest{
    return &AlibabaIdleParttimeNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeNotifyAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 实时同步类型, 0: 岗位状态, 1: 录取状态
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetType() int64 {
    return r._type
}
// Status Setter
// 岗位: 0关闭 ; 录取: 0不录取, 1录取
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetStatus() int64 {
    return r._status
}
// JobId Setter
// 岗位id
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetJobId(_jobId int64) error {
    r._jobId = _jobId
    r.Set("job_id", _jobId)
    return nil
}

// JobId Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetJobId() int64 {
    return r._jobId
}
// UserId Setter
// 用户id
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetUserId() int64 {
    return r._userId
}
// SyncTime Setter
// 同步时间
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetSyncTime(_syncTime int64) error {
    r._syncTime = _syncTime
    r.Set("sync_time", _syncTime)
    return nil
}

// SyncTime Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetSyncTime() int64 {
    return r._syncTime
}
// ApplyId Setter
// 报名id
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetApplyId() int64 {
    return r._applyId
}
// Message Setter
// 通知消息
func (r *AlibabaIdleParttimeNotifyAPIRequest) SetMessage(_message string) error {
    r._message = _message
    r.Set("message", _message)
    return nil
}

// Message Getter
func (r AlibabaIdleParttimeNotifyAPIRequest) GetMessage() string {
    return r._message
}