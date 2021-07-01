package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的电台 API请求
alibaba.xiami.api.radio.myself.get

我的电台
*/
type AlibabaXiamiApiRadioMyselfGetAPIRequest struct {
    model.Params
    // 歌曲数量
    _limit   int64
}

// 初始化AlibabaXiamiApiRadioMyselfGetAPIRequest对象
func NewAlibabaXiamiApiRadioMyselfGetRequest() *AlibabaXiamiApiRadioMyselfGetAPIRequest{
    return &AlibabaXiamiApiRadioMyselfGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.radio.myself.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 歌曲数量
func (r *AlibabaXiamiApiRadioMyselfGetAPIRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetLimit() int64 {
    return r._limit
}