package tbuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousersellergetAPIResponse 查询卖家用户信息 API返回值
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaousersellergetAPIResponse struct {
	model.CommonResponse
	TaobaousersellergetAPIResponseModel
}

// TaobaousersellergetAPIResponseModel is 查询卖家用户信息 成功返回结果
type TaobaousersellergetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_seller_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户
	User *User `json:"user,omitempty" xml:"user,omitempty"`
}