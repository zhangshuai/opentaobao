package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewPageuserroleAPIResponse 分页查询管理员 API返回值
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
type AlibabaCampusAclNewPageuserroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewPageuserroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewPageuserroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewPageuserroleAPIResponseModel).Reset()
}

// AlibabaCampusAclNewPageuserroleAPIResponseModel is 分页查询管理员 成功返回结果
type AlibabaCampusAclNewPageuserroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_pageuserrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewPageuserroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewPageuserroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewPageuserroleAPIResponse)
	},
}

// GetAlibabaCampusAclNewPageuserroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewPageuserroleAPIResponse
func GetAlibabaCampusAclNewPageuserroleAPIResponse() *AlibabaCampusAclNewPageuserroleAPIResponse {
	return poolAlibabaCampusAclNewPageuserroleAPIResponse.Get().(*AlibabaCampusAclNewPageuserroleAPIResponse)
}

// ReleaseAlibabaCampusAclNewPageuserroleAPIResponse 将 AlibabaCampusAclNewPageuserroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewPageuserroleAPIResponse(v *AlibabaCampusAclNewPageuserroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewPageuserroleAPIResponse.Put(v)
}
