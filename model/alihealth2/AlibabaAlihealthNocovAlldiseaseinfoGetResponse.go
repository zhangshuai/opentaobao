package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全国疫情统计数据 API返回值 
alibaba.alihealth.nocov.alldiseaseinfo.get

获取全国疫情统计数据
*/
type AlibabaAlihealthNocovAlldiseaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthNocovAlldiseaseinfoGetResponse
}

// 获取全国疫情统计数据 成功返回结果
type AlibabaAlihealthNocovAlldiseaseinfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_nocov_alldiseaseinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回全国疫情的统计数据
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // errCode
    BizErrCode   string `json:"biz_err_code,omitempty" xml:"biz_err_code,omitempty"`
    // errMessage
    BizErrMessage   string `json:"biz_err_message,omitempty" xml:"biz_err_message,omitempty"`
}