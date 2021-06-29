package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询订单标签 API请求
taobao.oc.tradetags.get

根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
*/
type TaobaoOcTradetagsGetRequest struct {
    model.Params
    // 交易主订单id
    tid   int64
    // 是否查询历史标签
    history   int64
    // 不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
    tagTypes   []string
    // 不填，则不做标签名称过滤
    tagNames   []string
}

// 初始化TaobaoOcTradetagsGetRequest对象
func NewTaobaoOcTradetagsGetRequest() *TaobaoOcTradetagsGetRequest{
    return &TaobaoOcTradetagsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetagsGetRequest) GetApiMethodName() string {
    return "taobao.oc.tradetags.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetagsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易主订单id
func (r *TaobaoOcTradetagsGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOcTradetagsGetRequest) GetTid() int64 {
    return r.tid
}
// History Setter
// 是否查询历史标签
func (r *TaobaoOcTradetagsGetRequest) SetHistory(history int64) error {
    r.history = history
    r.Set("history", history)
    return nil
}

// History Getter
func (r TaobaoOcTradetagsGetRequest) GetHistory() int64 {
    return r.history
}
// TagTypes Setter
// 不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
func (r *TaobaoOcTradetagsGetRequest) SetTagTypes(tagTypes []string) error {
    r.tagTypes = tagTypes
    r.Set("tag_types", tagTypes)
    return nil
}

// TagTypes Getter
func (r TaobaoOcTradetagsGetRequest) GetTagTypes() []string {
    return r.tagTypes
}
// TagNames Setter
// 不填，则不做标签名称过滤
func (r *TaobaoOcTradetagsGetRequest) SetTagNames(tagNames []string) error {
    r.tagNames = tagNames
    r.Set("tag_names", tagNames)
    return nil
}

// TagNames Getter
func (r TaobaoOcTradetagsGetRequest) GetTagNames() []string {
    return r.tagNames
}
