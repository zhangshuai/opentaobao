package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewpageuserroleAPIResponse 分页查询管理员 API返回值
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
type AlibabacampusaclnewpageuserroleAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewpageuserroleAPIResponseModel
}

// AlibabacampusaclnewpageuserroleAPIResponseModel is 分页查询管理员 成功返回结果
type AlibabacampusaclnewpageuserroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_pageuserrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
