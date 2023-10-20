package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyshopupdatestatusAPIResponse 更新渠道店营业状态 API返回值
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
type AlibabaaelophyshopupdatestatusAPIResponse struct {
	model.CommonResponse
	AlibabaaelophyshopupdatestatusAPIResponseModel
}

// AlibabaaelophyshopupdatestatusAPIResponseModel is 更新渠道店营业状态 成功返回结果
type AlibabaaelophyshopupdatestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_shop_updatestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api调用结果
	ApiResult *AlibabaaelophyshopupdatestatusApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}