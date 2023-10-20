package yunosdm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosdmsysgetdomainAPIResponse 获取动态域名 API返回值
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
type YunosdmsysgetdomainAPIResponse struct {
	model.CommonResponse
	YunosdmsysgetdomainAPIResponseModel
}

// YunosdmsysgetdomainAPIResponseModel is 获取动态域名 成功返回结果
type YunosdmsysgetdomainAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_dm_sys_get_domain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// obj
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
