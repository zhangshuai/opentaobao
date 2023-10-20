package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdmodifyAPIResponse 覆盖单元下同类型定向人群 API返回值
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
type TaobaofeedflowitemcrowdmodifyAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcrowdmodifyAPIResponseModel
}

// TaobaofeedflowitemcrowdmodifyAPIResponseModel is 覆盖单元下同类型定向人群 成功返回结果
type TaobaofeedflowitemcrowdmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaofeedflowitemcrowdmodifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
