package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRoomtypeStatusUpdateAPIResponse top房型状态修改 API返回值
// taobao.roomtype.status.update
//
// top房型状态修改
type TaobaoRoomtypeStatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRoomtypeStatusUpdateAPIResponseModel
}

// TaobaoRoomtypeStatusUpdateAPIResponseModel is top房型状态修改 成功返回结果
type TaobaoRoomtypeStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"roomtype_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
