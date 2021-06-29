package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询子订单对应的评价、追评以及语义标签 API返回值 
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
*/
type TmallTraderateFeedsGetAPIResponse struct {
    model.CommonResponse
    TmallTraderateFeedsGetResponse
}

// 查询子订单对应的评价、追评以及语义标签 成功返回结果
type TmallTraderateFeedsGetResponse struct {
    XMLName xml.Name `xml:"tmall_traderate_feeds_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回评价信息
    TmallRateInfo   *TmallRateInfo `json:"tmall_rate_info,omitempty" xml:"tmall_rate_info,omitempty"`
}
