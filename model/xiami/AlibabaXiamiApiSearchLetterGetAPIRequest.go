package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索接口（首字母） API请求
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
type AlibabaXiamiApiSearchLetterGetAPIRequest struct {
    model.Params
    // 搜索关键字
    _key   string
}

// 初始化AlibabaXiamiApiSearchLetterGetAPIRequest对象
func NewAlibabaXiamiApiSearchLetterGetRequest() *AlibabaXiamiApiSearchLetterGetAPIRequest{
    return &AlibabaXiamiApiSearchLetterGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiSearchLetterGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.letter.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiSearchLetterGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Key Setter
// 搜索关键字
func (r *AlibabaXiamiApiSearchLetterGetAPIRequest) SetKey(_key string) error {
    r._key = _key
    r.Set("key", _key)
    return nil
}

// Key Getter
func (r AlibabaXiamiApiSearchLetterGetAPIRequest) GetKey() string {
    return r._key
}