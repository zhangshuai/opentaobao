package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductdownshelfAPIResponse 产品下架 API返回值
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
type TmallsupplychainchannelproductdownshelfAPIResponse struct {
	model.CommonResponse
	TmallsupplychainchannelproductdownshelfAPIResponseModel
}

// TmallsupplychainchannelproductdownshelfAPIResponseModel is 产品下架 成功返回结果
type TmallsupplychainchannelproductdownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallsupplychainchannelproductdownshelfResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
