package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliverysellchannelconfigsqueryAPIResponse 查询商家配置的信息 API返回值
// taobao.miniapp.ext.delivery.sell.channel.configs.query
//
// 查询商家配置的信息
type TaobaominiappextdeliverysellchannelconfigsqueryAPIResponse struct {
	model.CommonResponse
	TaobaominiappextdeliverysellchannelconfigsqueryAPIResponseModel
}

// TaobaominiappextdeliverysellchannelconfigsqueryAPIResponseModel is 查询商家配置的信息 成功返回结果
type TaobaominiappextdeliverysellchannelconfigsqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_sell_channel_configs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MiniappResult `json:"result,omitempty" xml:"result,omitempty"`
}