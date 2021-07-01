package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资质查询接口 API请求
taobao.tanx.qualification.find

资质查询接口
*/
type TaobaoTanxQualificationFindAPIRequest struct {
    model.Params
    // dsp客户在tanx的memberId
    _memberId   int64
    // dsp客户在tanx签名的token
    _token   string
    // 当前时间,从1970年算的毫秒
    _signTime   int64
    // 分页的第几页，从1开始
    _page   int64
    // 分页大小，限制200
    _pageSize   int64
    // 广告主资质查询dto
    _query   *QualificationQuery
}

// 初始化TaobaoTanxQualificationFindAPIRequest对象
func NewTaobaoTanxQualificationFindRequest() *TaobaoTanxQualificationFindAPIRequest{
    return &TaobaoTanxQualificationFindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationFindAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationFindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// dsp客户在tanx的memberId
func (r *TaobaoTanxQualificationFindAPIRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp客户在tanx签名的token
func (r *TaobaoTanxQualificationFindAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间,从1970年算的毫秒
func (r *TaobaoTanxQualificationFindAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetSignTime() int64 {
    return r._signTime
}
// Page Setter
// 分页的第几页，从1开始
func (r *TaobaoTanxQualificationFindAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 分页大小，限制200
func (r *TaobaoTanxQualificationFindAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Query Setter
// 广告主资质查询dto
func (r *TaobaoTanxQualificationFindAPIRequest) SetQuery(_query *QualificationQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoTanxQualificationFindAPIRequest) GetQuery() *QualificationQuery {
    return r._query
}