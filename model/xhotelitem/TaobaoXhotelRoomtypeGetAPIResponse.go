package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypegetAPIResponse 房型查询接口 API返回值
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
type TaobaoxhotelroomtypegetAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelroomtypegetAPIResponseModel
}

// TaobaoxhotelroomtypegetAPIResponseModel is 房型查询接口 成功返回结果
type TaobaoxhotelroomtypegetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询得到的RoomType
	Xroomtype *XroomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
