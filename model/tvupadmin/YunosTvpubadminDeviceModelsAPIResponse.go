package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicemodelsAPIResponse 获取品牌下设备列表 API返回值
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
type YunostvpubadmindevicemodelsAPIResponse struct {
	model.CommonResponse
	YunostvpubadmindevicemodelsAPIResponseModel
}

// YunostvpubadmindevicemodelsAPIResponseModel is 获取品牌下设备列表 成功返回结果
type YunostvpubadmindevicemodelsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_models_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	ModelList []string `json:"model_list,omitempty" xml:"model_list>string,omitempty"`
}
