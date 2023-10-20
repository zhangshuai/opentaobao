package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweitaofeedsynchronizeAPIResponse 推广淘小铺isv 活动到微淘feed API返回值
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
type TaobaoweitaofeedsynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoweitaofeedsynchronizeAPIResponseModel
}

// TaobaoweitaofeedsynchronizeAPIResponseModel is 推广淘小铺isv 活动到微淘feed 成功返回结果
type TaobaoweitaofeedsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_feed_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 增加错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 同步到微淘成功与否
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
