package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelatesubgetAPIResponse 门店和子门店关系查找 API返回值
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
type TaobaoplacestorerelatesubgetAPIResponse struct {
	model.CommonResponse
	TaobaoplacestorerelatesubgetAPIResponseModel
}

// TaobaoplacestorerelatesubgetAPIResponseModel is 门店和子门店关系查找 成功返回结果
type TaobaoplacestorerelatesubgetAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storerelatesub_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
