package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemrechargebatchremoveAPIResponse 闲鱼商品直充功能移除 API返回值
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
type AlibabaidleisvitemrechargebatchremoveAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvitemrechargebatchremoveAPIResponseModel
}

// AlibabaidleisvitemrechargebatchremoveAPIResponseModel is 闲鱼商品直充功能移除 成功返回结果
type AlibabaidleisvitemrechargebatchremoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_recharge_batch_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TopListResult `json:"result,omitempty" xml:"result,omitempty"`
}