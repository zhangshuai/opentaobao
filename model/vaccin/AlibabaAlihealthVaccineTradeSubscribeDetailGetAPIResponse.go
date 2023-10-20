package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccinetradesubscribedetailgetAPIResponse 私立疫苗交易-预约详情获取 API返回值
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
type AlibabaalihealthvaccinetradesubscribedetailgetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthvaccinetradesubscribedetailgetAPIResponseModel
}

// AlibabaalihealthvaccinetradesubscribedetailgetAPIResponseModel is 私立疫苗交易-预约详情获取 成功返回结果
type AlibabaalihealthvaccinetradesubscribedetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_trade_subscribe_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务报错code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务报错消息
	BizMessage string `json:"biz_message,omitempty" xml:"biz_message,omitempty"`
	// 预约详情响应
	Data *TradeVaccineSubscribeDetailTopResult `json:"data,omitempty" xml:"data,omitempty"`
	// 业务成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
