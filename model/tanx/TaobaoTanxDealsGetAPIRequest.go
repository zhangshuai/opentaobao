package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取交易列表 API请求
taobao.tanx.deals.get

批量获取交易信息
*/
type TaobaoTanxDealsGetAPIRequest struct {
    model.Params
    // dsp用户id
    _dspId   int64
    // dsp用户验证token
    _token   string
    // 页大小
    _pageSize   int64
    // 交易类型
    _dealType   int64
    // 页码
    _page   int64
    // 1970年到现在的时间，毫秒
    _signTime   int64
}

// 初始化TaobaoTanxDealsGetAPIRequest对象
func NewTaobaoTanxDealsGetRequest() *TaobaoTanxDealsGetAPIRequest{
    return &TaobaoTanxDealsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxDealsGetAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.deals.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxDealsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DspId Setter
// dsp用户id
func (r *TaobaoTanxDealsGetAPIRequest) SetDspId(_dspId int64) error {
    r._dspId = _dspId
    r.Set("dsp_id", _dspId)
    return nil
}

// DspId Getter
func (r TaobaoTanxDealsGetAPIRequest) GetDspId() int64 {
    return r._dspId
}
// Token Setter
// dsp用户验证token
func (r *TaobaoTanxDealsGetAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxDealsGetAPIRequest) GetToken() string {
    return r._token
}
// PageSize Setter
// 页大小
func (r *TaobaoTanxDealsGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTanxDealsGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// DealType Setter
// 交易类型
func (r *TaobaoTanxDealsGetAPIRequest) SetDealType(_dealType int64) error {
    r._dealType = _dealType
    r.Set("deal_type", _dealType)
    return nil
}

// DealType Getter
func (r TaobaoTanxDealsGetAPIRequest) GetDealType() int64 {
    return r._dealType
}
// Page Setter
// 页码
func (r *TaobaoTanxDealsGetAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoTanxDealsGetAPIRequest) GetPage() int64 {
    return r._page
}
// SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaoTanxDealsGetAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxDealsGetAPIRequest) GetSignTime() int64 {
    return r._signTime
}