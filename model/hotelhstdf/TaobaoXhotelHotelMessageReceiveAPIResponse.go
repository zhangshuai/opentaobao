package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelhotelmessagereceiveAPIResponse 接收道消息接口 API返回值
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
type TaobaoxhotelhotelmessagereceiveAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelhotelmessagereceiveAPIResponseModel
}

// TaobaoxhotelhotelmessagereceiveAPIResponseModel is 接收道消息接口 成功返回结果
type TaobaoxhotelhotelmessagereceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_hotel_message_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *DidaResult `json:"result,omitempty" xml:"result,omitempty"`
}