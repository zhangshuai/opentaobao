package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemanchorgetAPIResponse 获取可用宝贝描述规范化模块 API返回值
// taobao.item.anchor.get
//
// 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoitemanchorgetAPIResponse struct {
	model.CommonResponse
	TaobaoitemanchorgetAPIResponseModel
}

// TaobaoitemanchorgetAPIResponseModel is 获取可用宝贝描述规范化模块 成功返回结果
type TaobaoitemanchorgetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_anchor_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 宝贝描述规范化可使用打标模块的锚点信息
	AnchorModules []IdsModule `json:"anchor_modules,omitempty" xml:"anchor_modules>ids_module,omitempty"`
	// 返回的宝贝描述模板结果数目
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}