package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturegetAPIResponse 获取图片信息 API返回值
// taobao.picture.get
//
// 获取图片信息
type TaobaopicturegetAPIResponse struct {
	model.CommonResponse
	TaobaopicturegetAPIResponseModel
}

// TaobaopicturegetAPIResponseModel is 获取图片信息 成功返回结果
type TaobaopicturegetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息列表
	Pictures []Picture `json:"pictures,omitempty" xml:"pictures>picture,omitempty"`
}
