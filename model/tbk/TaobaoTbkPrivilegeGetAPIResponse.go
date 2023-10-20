package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkprivilegegetAPIResponse 淘宝客-服务商-单品券高效转链 API返回值
// taobao.tbk.privilege.get
//
// 单品券高效转链API
type TaobaotbkprivilegegetAPIResponse struct {
	model.CommonResponse
	TaobaotbkprivilegegetAPIResponseModel
}

// TaobaotbkprivilegegetAPIResponseModel is 淘宝客-服务商-单品券高效转链 成功返回结果
type TaobaotbkprivilegegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_privilege_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaotbkprivilegegetRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
