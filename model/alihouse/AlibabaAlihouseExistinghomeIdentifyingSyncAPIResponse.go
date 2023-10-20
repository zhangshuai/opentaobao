package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeidentifyingsyncAPIResponse 登陆标识信息同步 API返回值
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
type AlibabaalihouseexistinghomeidentifyingsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeidentifyingsyncAPIResponseModel
}

// AlibabaalihouseexistinghomeidentifyingsyncAPIResponseModel is 登陆标识信息同步 成功返回结果
type AlibabaalihouseexistinghomeidentifyingsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_identifying_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回信息
	Result *AlibabaalihouseexistinghomeidentifyingsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}