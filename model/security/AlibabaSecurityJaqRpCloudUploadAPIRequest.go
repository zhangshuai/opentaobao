package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云上传接口 API请求
alibaba.security.jaq.rp.cloud.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpCloudUploadAPIRequest struct {
    model.Params
    // 认证token
    _verifyToken   string
    // []
    _elements   []Elements
}

// 初始化AlibabaSecurityJaqRpCloudUploadAPIRequest对象
func NewAlibabaSecurityJaqRpCloudUploadRequest() *AlibabaSecurityJaqRpCloudUploadAPIRequest{
    return &AlibabaSecurityJaqRpCloudUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudUploadAPIRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetVerifyToken() string {
    return r._verifyToken
}
// Elements Setter
// []
func (r *AlibabaSecurityJaqRpCloudUploadAPIRequest) SetElements(_elements []Elements) error {
    r._elements = _elements
    r.Set("elements", _elements)
    return nil
}

// Elements Getter
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetElements() []Elements {
    return r._elements
}