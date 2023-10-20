package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryUpdateAPIResponse 后端商品库存增量更新接口 API返回值
// taobao.jst.astrolabe.storeinventory.update
//
// 增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存
type TaobaoJstAstrolabeStoreinventoryUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryUpdateAPIResponseModel
}

// TaobaoJstAstrolabeStoreinventoryUpdateAPIResponseModel is 后端商品库存增量更新接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息列表
	ErrorDescriptions []TaobaoJstAstrolabeStoreinventoryUpdateError `json:"error_descriptions,omitempty" xml:"error_descriptions>taobao_jst_astrolabe_storeinventory_update_error,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应标签
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
