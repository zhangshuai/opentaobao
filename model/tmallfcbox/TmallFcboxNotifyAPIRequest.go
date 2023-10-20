package tmallfcbox

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFcboxNotifyAPIRequest 丰巢通知接口 API请求
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
type TmallFcboxNotifyAPIRequest struct {
	model.Params
	// 状态备注
	_stateRemark string
	// 申请记录当前状态
	_state string
	// 变更时间
	_operateTime string
	// 变更操作
	_operate string
	// 申请接口返回的申请标识
	_applyId string
}

// NewTmallFcboxNotifyRequest 初始化TmallFcboxNotifyAPIRequest对象
func NewTmallFcboxNotifyRequest() *TmallFcboxNotifyAPIRequest {
	return &TmallFcboxNotifyAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFcboxNotifyAPIRequest) Reset() {
	r._stateRemark = ""
	r._state = ""
	r._operateTime = ""
	r._operate = ""
	r._applyId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFcboxNotifyAPIRequest) GetApiMethodName() string {
	return "tmall.fcbox.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFcboxNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFcboxNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStateRemark is StateRemark Setter
// 状态备注
func (r *TmallFcboxNotifyAPIRequest) SetStateRemark(_stateRemark string) error {
	r._stateRemark = _stateRemark
	r.Set("state_remark", _stateRemark)
	return nil
}

// GetStateRemark StateRemark Getter
func (r TmallFcboxNotifyAPIRequest) GetStateRemark() string {
	return r._stateRemark
}

// SetState is State Setter
// 申请记录当前状态
func (r *TmallFcboxNotifyAPIRequest) SetState(_state string) error {
	r._state = _state
	r.Set("state", _state)
	return nil
}

// GetState State Getter
func (r TmallFcboxNotifyAPIRequest) GetState() string {
	return r._state
}

// SetOperateTime is OperateTime Setter
// 变更时间
func (r *TmallFcboxNotifyAPIRequest) SetOperateTime(_operateTime string) error {
	r._operateTime = _operateTime
	r.Set("operate_time", _operateTime)
	return nil
}

// GetOperateTime OperateTime Getter
func (r TmallFcboxNotifyAPIRequest) GetOperateTime() string {
	return r._operateTime
}

// SetOperate is Operate Setter
// 变更操作
func (r *TmallFcboxNotifyAPIRequest) SetOperate(_operate string) error {
	r._operate = _operate
	r.Set("operate", _operate)
	return nil
}

// GetOperate Operate Getter
func (r TmallFcboxNotifyAPIRequest) GetOperate() string {
	return r._operate
}

// SetApplyId is ApplyId Setter
// 申请接口返回的申请标识
func (r *TmallFcboxNotifyAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TmallFcboxNotifyAPIRequest) GetApplyId() string {
	return r._applyId
}

var poolTmallFcboxNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFcboxNotifyRequest()
	},
}

// GetTmallFcboxNotifyRequest 从 sync.Pool 获取 TmallFcboxNotifyAPIRequest
func GetTmallFcboxNotifyAPIRequest() *TmallFcboxNotifyAPIRequest {
	return poolTmallFcboxNotifyAPIRequest.Get().(*TmallFcboxNotifyAPIRequest)
}

// ReleaseTmallFcboxNotifyAPIRequest 将 TmallFcboxNotifyAPIRequest 放入 sync.Pool
func ReleaseTmallFcboxNotifyAPIRequest(v *TmallFcboxNotifyAPIRequest) {
	v.Reset()
	poolTmallFcboxNotifyAPIRequest.Put(v)
}
