package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoreupdateAPIResponse 门店更新接口 API返回值
// taobao.qimen.store.update
//
// 商家在ERP等系统中调用该接口，更新门店信息
type TaobaoqimenstoreupdateAPIResponse struct {
	model.CommonResponse
	TaobaoqimenstoreupdateAPIResponseModel
}

// TaobaoqimenstoreupdateAPIResponseModel is 门店更新接口 成功返回结果
type TaobaoqimenstoreupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
