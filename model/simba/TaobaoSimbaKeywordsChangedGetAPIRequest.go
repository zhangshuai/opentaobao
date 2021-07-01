package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的关键词ID、宝贝id、修改时间 API请求
taobao.simba.keywords.changed.get

分页获取修改过的关键词ID、宝贝id、修改时间
*/
type TaobaoSimbaKeywordsChangedGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 得到此时间点之后的数据，不能大于一个月
    _startTime   string
    // 返回的每页数据量大小,默认200最大1000
    _pageSize   int64
    // 返回的第几页数据，默认为1
    _pageNo   int64
}

// 初始化TaobaoSimbaKeywordsChangedGetAPIRequest对象
func NewTaobaoSimbaKeywordsChangedGetRequest() *TaobaoSimbaKeywordsChangedGetAPIRequest{
    return &TaobaoSimbaKeywordsChangedGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.changed.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsChangedGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaKeywordsChangedGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaKeywordsChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordsChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaKeywordsChangedGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}