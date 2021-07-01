package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活风控数据接入 API请求
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口
*/
type AlibabaAlscKmsAccessAPIRequest struct {
    model.Params
    // 接入参数
    _requestdata   string
}

// 初始化AlibabaAlscKmsAccessAPIRequest对象
func NewAlibabaAlscKmsAccessRequest() *AlibabaAlscKmsAccessAPIRequest{
    return &AlibabaAlscKmsAccessAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscKmsAccessAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.kms.access"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscKmsAccessAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Requestdata Setter
// 接入参数
func (r *AlibabaAlscKmsAccessAPIRequest) SetRequestdata(_requestdata string) error {
    r._requestdata = _requestdata
    r.Set("requestdata", _requestdata)
    return nil
}

// Requestdata Getter
func (r AlibabaAlscKmsAccessAPIRequest) GetRequestdata() string {
    return r._requestdata
}