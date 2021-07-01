package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除创意 API请求
taobao.simba.creative.delete

删除一个创意
*/
type TaobaoSimbaCreativeDeleteAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 创意Id
    _creativeId   int64
}

// 初始化TaobaoSimbaCreativeDeleteAPIRequest对象
func NewTaobaoSimbaCreativeDeleteRequest() *TaobaoSimbaCreativeDeleteAPIRequest{
    return &TaobaoSimbaCreativeDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.simba.creative.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeDeleteAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetNick() string {
    return r._nick
}
// CreativeId Setter
// 创意Id
func (r *TaobaoSimbaCreativeDeleteAPIRequest) SetCreativeId(_creativeId int64) error {
    r._creativeId = _creativeId
    r.Set("creative_id", _creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetCreativeId() int64 {
    return r._creativeId
}