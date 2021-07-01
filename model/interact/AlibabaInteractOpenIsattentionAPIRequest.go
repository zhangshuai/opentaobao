package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户是否收藏某个店铺 API请求
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
type AlibabaInteractOpenIsattentionAPIRequest struct {
    model.Params
    // 1
    _param0   int64
}

// 初始化AlibabaInteractOpenIsattentionAPIRequest对象
func NewAlibabaInteractOpenIsattentionRequest() *AlibabaInteractOpenIsattentionAPIRequest{
    return &AlibabaInteractOpenIsattentionAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractOpenIsattentionAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.open.isattention"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractOpenIsattentionAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 1
func (r *AlibabaInteractOpenIsattentionAPIRequest) SetParam0(_param0 int64) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaInteractOpenIsattentionAPIRequest) GetParam0() int64 {
    return r._param0
}