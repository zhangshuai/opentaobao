package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证上传认证信息 APIResponse
alibaba.security.jaq.rp.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqRpUploadResponse `json:"alibaba_security_jaq_rp_upload_response,omitempty"`
}

type AlibabaSecurityJaqRpUploadResponse struct {

    // 返回信息
    Data   *RpUploadResult `json:"data,omitempty"`

}