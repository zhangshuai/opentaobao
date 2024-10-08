package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationFindAPIRequest 资质查询接口 API请求
// taobao.tanx.qualification.find
//
// 资质查询接口
type TaobaoTanxQualificationFindAPIRequest struct {
	model.Params
	// dsp客户在tanx签名的token
	_token string
	// dsp客户在tanx的memberId
	_memberId int64
	// 当前时间,从1970年算的毫秒
	_signTime int64
	// 分页的第几页，从1开始
	_page int64
	// 分页大小，限制200
	_pageSize int64
	// 广告主资质查询dto
	_query *QualificationQuery
}

// NewTaobaoTanxQualificationFindRequest 初始化TaobaoTanxQualificationFindAPIRequest对象
func NewTaobaoTanxQualificationFindRequest() *TaobaoTanxQualificationFindAPIRequest {
	return &TaobaoTanxQualificationFindAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxQualificationFindAPIRequest) Reset() {
	r._token = ""
	r._memberId = 0
	r._signTime = 0
	r._page = 0
	r._pageSize = 0
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationFindAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxQualificationFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp客户在tanx签名的token
func (r *TaobaoTanxQualificationFindAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// dsp客户在tanx的memberId
func (r *TaobaoTanxQualificationFindAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间,从1970年算的毫秒
func (r *TaobaoTanxQualificationFindAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 分页的第几页，从1开始
func (r *TaobaoTanxQualificationFindAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 分页大小，限制200
func (r *TaobaoTanxQualificationFindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetQuery is Query Setter
// 广告主资质查询dto
func (r *TaobaoTanxQualificationFindAPIRequest) SetQuery(_query *QualificationQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetQuery() *QualificationQuery {
	return r._query
}

var poolTaobaoTanxQualificationFindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxQualificationFindRequest()
	},
}

// GetTaobaoTanxQualificationFindRequest 从 sync.Pool 获取 TaobaoTanxQualificationFindAPIRequest
func GetTaobaoTanxQualificationFindAPIRequest() *TaobaoTanxQualificationFindAPIRequest {
	return poolTaobaoTanxQualificationFindAPIRequest.Get().(*TaobaoTanxQualificationFindAPIRequest)
}

// ReleaseTaobaoTanxQualificationFindAPIRequest 将 TaobaoTanxQualificationFindAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxQualificationFindAPIRequest(v *TaobaoTanxQualificationFindAPIRequest) {
	v.Reset()
	poolTaobaoTanxQualificationFindAPIRequest.Put(v)
}
