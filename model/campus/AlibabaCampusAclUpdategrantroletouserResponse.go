package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改用户到角色关系 API返回值 
alibaba.campus.acl.updategrantroletouser

修改用户到角色关系
*/
type AlibabaCampusAclUpdategrantroletouserAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclUpdategrantroletouserResponse
}

// 修改用户到角色关系 成功返回结果
type AlibabaCampusAclUpdategrantroletouserResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_updategrantroletouser_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}