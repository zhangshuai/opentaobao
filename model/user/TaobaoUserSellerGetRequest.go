package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家用户信息 API请求
taobao.user.seller.get

查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
*/
type TaobaoUserSellerGetRequest struct {
    model.Params
    // 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
    fields   []string
}

// 初始化TaobaoUserSellerGetRequest对象
func NewTaobaoUserSellerGetRequest() *TaobaoUserSellerGetRequest{
    return &TaobaoUserSellerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUserSellerGetRequest) GetApiMethodName() string {
    return "taobao.user.seller.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUserSellerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
func (r *TaobaoUserSellerGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoUserSellerGetRequest) GetFields() []string {
    return r.fields
}
