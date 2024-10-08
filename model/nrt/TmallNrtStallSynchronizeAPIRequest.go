package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStallSynchronizeAPIRequest 摊位信息同步 API请求
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
type TmallNrtStallSynchronizeAPIRequest struct {
	model.Params
	// 参数对象
	_stall *NrtStoreDto
}

// NewTmallNrtStallSynchronizeRequest 初始化TmallNrtStallSynchronizeAPIRequest对象
func NewTmallNrtStallSynchronizeRequest() *TmallNrtStallSynchronizeAPIRequest {
	return &TmallNrtStallSynchronizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtStallSynchronizeAPIRequest) Reset() {
	r._stall = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStallSynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.stall.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStallSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStallSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStall is Stall Setter
// 参数对象
func (r *TmallNrtStallSynchronizeAPIRequest) SetStall(_stall *NrtStoreDto) error {
	r._stall = _stall
	r.Set("stall", _stall)
	return nil
}

// GetStall Stall Getter
func (r TmallNrtStallSynchronizeAPIRequest) GetStall() *NrtStoreDto {
	return r._stall
}

var poolTmallNrtStallSynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtStallSynchronizeRequest()
	},
}

// GetTmallNrtStallSynchronizeRequest 从 sync.Pool 获取 TmallNrtStallSynchronizeAPIRequest
func GetTmallNrtStallSynchronizeAPIRequest() *TmallNrtStallSynchronizeAPIRequest {
	return poolTmallNrtStallSynchronizeAPIRequest.Get().(*TmallNrtStallSynchronizeAPIRequest)
}

// ReleaseTmallNrtStallSynchronizeAPIRequest 将 TmallNrtStallSynchronizeAPIRequest 放入 sync.Pool
func ReleaseTmallNrtStallSynchronizeAPIRequest(v *TmallNrtStallSynchronizeAPIRequest) {
	v.Reset()
	poolTmallNrtStallSynchronizeAPIRequest.Put(v)
}
