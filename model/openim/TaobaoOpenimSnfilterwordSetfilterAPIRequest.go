package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimSnfilterwordSetfilterAPIRequest 关键词过滤 API请求
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
type TaobaoOpenimSnfilterwordSetfilterAPIRequest struct {
	model.Params
	// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
	_creator string
	// 需要过滤的关键词
	_filterword string
	// 过滤原因描述
	_desc string
}

// NewTaobaoOpenimSnfilterwordSetfilterRequest 初始化TaobaoOpenimSnfilterwordSetfilterAPIRequest对象
func NewTaobaoOpenimSnfilterwordSetfilterRequest() *TaobaoOpenimSnfilterwordSetfilterAPIRequest {
	return &TaobaoOpenimSnfilterwordSetfilterAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) Reset() {
	r._creator = ""
	r._filterword = ""
	r._desc = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetApiMethodName() string {
	return "taobao.openim.snfilterword.setfilter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreator is Creator Setter
// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetCreator(_creator string) error {
	r._creator = _creator
	r.Set("creator", _creator)
	return nil
}

// GetCreator Creator Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetCreator() string {
	return r._creator
}

// SetFilterword is Filterword Setter
// 需要过滤的关键词
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetFilterword(_filterword string) error {
	r._filterword = _filterword
	r.Set("filterword", _filterword)
	return nil
}

// GetFilterword Filterword Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetFilterword() string {
	return r._filterword
}

// SetDesc is Desc Setter
// 过滤原因描述
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetDesc() string {
	return r._desc
}

var poolTaobaoOpenimSnfilterwordSetfilterAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimSnfilterwordSetfilterRequest()
	},
}

// GetTaobaoOpenimSnfilterwordSetfilterRequest 从 sync.Pool 获取 TaobaoOpenimSnfilterwordSetfilterAPIRequest
func GetTaobaoOpenimSnfilterwordSetfilterAPIRequest() *TaobaoOpenimSnfilterwordSetfilterAPIRequest {
	return poolTaobaoOpenimSnfilterwordSetfilterAPIRequest.Get().(*TaobaoOpenimSnfilterwordSetfilterAPIRequest)
}

// ReleaseTaobaoOpenimSnfilterwordSetfilterAPIRequest 将 TaobaoOpenimSnfilterwordSetfilterAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimSnfilterwordSetfilterAPIRequest(v *TaobaoOpenimSnfilterwordSetfilterAPIRequest) {
	v.Reset()
	poolTaobaoOpenimSnfilterwordSetfilterAPIRequest.Put(v)
}
