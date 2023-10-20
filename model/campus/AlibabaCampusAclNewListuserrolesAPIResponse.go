package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistuserrolesAPIResponse 查询并标记用户选择的角色 API返回值
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
type AlibabacampusaclnewlistuserrolesAPIResponse struct {
	model.CommonResponse
	AlibabacampusaclnewlistuserrolesAPIResponseModel
}

// AlibabacampusaclnewlistuserrolesAPIResponseModel is 查询并标记用户选择的角色 成功返回结果
type AlibabacampusaclnewlistuserrolesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listuserroles_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
