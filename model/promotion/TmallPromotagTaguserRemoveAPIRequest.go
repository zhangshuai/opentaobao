package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户移除优惠标签 API请求
tmall.promotag.taguser.remove

给用户载体去标
*/
type TmallPromotagTaguserRemoveAPIRequest struct {
    model.Params
    // 标签ID
    _tagId   int64
    // 买家昵称
    _nick   string
}

// 初始化TmallPromotagTaguserRemoveAPIRequest对象
func NewTmallPromotagTaguserRemoveRequest() *TmallPromotagTaguserRemoveAPIRequest{
    return &TmallPromotagTaguserRemoveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserRemoveAPIRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.remove"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserRemoveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTaguserRemoveAPIRequest) SetTagId(_tagId int64) error {
    r._tagId = _tagId
    r.Set("tag_id", _tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetTagId() int64 {
    return r._tagId
}
// Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserRemoveAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetNick() string {
    return r._nick
}