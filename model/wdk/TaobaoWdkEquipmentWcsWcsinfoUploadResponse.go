package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链业务信息上传 API返回值 
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentWcsWcsinfoUploadResponse
}

// 悬挂链业务信息上传 成功返回结果
type TaobaoWdkEquipmentWcsWcsinfoUploadResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_wcs_wcsinfo_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // model
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
