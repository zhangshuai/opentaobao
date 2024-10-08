package pentraprism

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPentaprismTaskTriggerAPIRequest 推进单条任务进度 API请求
// taobao.pentaprism.task.trigger
//
// 外网用户推进单条五棱镜任务进度
type TaobaoPentaprismTaskTriggerAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaoPentaprismTaskTriggerRequest 初始化TaobaoPentaprismTaskTriggerAPIRequest对象
func NewTaobaoPentaprismTaskTriggerRequest() *TaobaoPentaprismTaskTriggerAPIRequest {
	return &TaobaoPentaprismTaskTriggerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPentaprismTaskTriggerAPIRequest) Reset() {
	r._openPo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.trigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenPo is OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskTriggerAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// GetOpenPo OpenPo Getter
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}

var poolTaobaoPentaprismTaskTriggerAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPentaprismTaskTriggerRequest()
	},
}

// GetTaobaoPentaprismTaskTriggerRequest 从 sync.Pool 获取 TaobaoPentaprismTaskTriggerAPIRequest
func GetTaobaoPentaprismTaskTriggerAPIRequest() *TaobaoPentaprismTaskTriggerAPIRequest {
	return poolTaobaoPentaprismTaskTriggerAPIRequest.Get().(*TaobaoPentaprismTaskTriggerAPIRequest)
}

// ReleaseTaobaoPentaprismTaskTriggerAPIRequest 将 TaobaoPentaprismTaskTriggerAPIRequest 放入 sync.Pool
func ReleaseTaobaoPentaprismTaskTriggerAPIRequest(v *TaobaoPentaprismTaskTriggerAPIRequest) {
	v.Reset()
	poolTaobaoPentaprismTaskTriggerAPIRequest.Put(v)
}
