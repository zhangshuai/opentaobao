package wangwang

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuKefuevalGetAPIRequest 客服评价详情接口 API请求
// taobao.qianniu.kefueval.get
//
// 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetAPIRequest struct {
	model.Params
	// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
	_queryIds string
	// 开始时间，格式yyyyMMddHHmmss
	_btime string
	// 结束时间，格式yyyyMMddHHmmss
	_etime string
}

// NewTaobaoQianniuKefuevalGetRequest 初始化TaobaoQianniuKefuevalGetAPIRequest对象
func NewTaobaoQianniuKefuevalGetRequest() *TaobaoQianniuKefuevalGetAPIRequest {
	return &TaobaoQianniuKefuevalGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuKefuevalGetAPIRequest) Reset() {
	r._queryIds = ""
	r._btime = ""
	r._etime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuKefuevalGetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.kefueval.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuKefuevalGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuKefuevalGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryIds is QueryIds Setter
// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetQueryIds(_queryIds string) error {
	r._queryIds = _queryIds
	r.Set("query_ids", _queryIds)
	return nil
}

// GetQueryIds QueryIds Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetQueryIds() string {
	return r._queryIds
}

// SetBtime is Btime Setter
// 开始时间，格式yyyyMMddHHmmss
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetBtime(_btime string) error {
	r._btime = _btime
	r.Set("btime", _btime)
	return nil
}

// GetBtime Btime Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetBtime() string {
	return r._btime
}

// SetEtime is Etime Setter
// 结束时间，格式yyyyMMddHHmmss
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// GetEtime Etime Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetEtime() string {
	return r._etime
}

var poolTaobaoQianniuKefuevalGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuKefuevalGetRequest()
	},
}

// GetTaobaoQianniuKefuevalGetRequest 从 sync.Pool 获取 TaobaoQianniuKefuevalGetAPIRequest
func GetTaobaoQianniuKefuevalGetAPIRequest() *TaobaoQianniuKefuevalGetAPIRequest {
	return poolTaobaoQianniuKefuevalGetAPIRequest.Get().(*TaobaoQianniuKefuevalGetAPIRequest)
}

// ReleaseTaobaoQianniuKefuevalGetAPIRequest 将 TaobaoQianniuKefuevalGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuKefuevalGetAPIRequest(v *TaobaoQianniuKefuevalGetAPIRequest) {
	v.Reset()
	poolTaobaoQianniuKefuevalGetAPIRequest.Put(v)
}
