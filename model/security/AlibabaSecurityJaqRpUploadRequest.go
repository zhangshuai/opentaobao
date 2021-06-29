package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证上传认证信息 API请求
alibaba.security.jaq.rp.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpUploadRequest struct {
    model.Params
    // 认证会话token
    verifyToken   string
    // 此次需要上传的认证信息的列表
    elements   []Element
}

// 初始化AlibabaSecurityJaqRpUploadRequest对象
func NewAlibabaSecurityJaqRpUploadRequest() *AlibabaSecurityJaqRpUploadRequest{
    return &AlibabaSecurityJaqRpUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpUploadRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证会话token
func (r *AlibabaSecurityJaqRpUploadRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpUploadRequest) GetVerifyToken() string {
    return r.verifyToken
}
// Elements Setter
// 此次需要上传的认证信息的列表
func (r *AlibabaSecurityJaqRpUploadRequest) SetElements(elements []Element) error {
    r.elements = elements
    r.Set("elements", elements)
    return nil
}

// Elements Getter
func (r AlibabaSecurityJaqRpUploadRequest) GetElements() []Element {
    return r.elements
}
