package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradessoldqueryAPIResponse 根据收件人信息查询交易单号 API返回值
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
type TaobaotradessoldqueryAPIResponse struct {
	model.CommonResponse
	TaobaotradessoldqueryAPIResponseModel
}

// TaobaotradessoldqueryAPIResponseModel is 根据收件人信息查询交易单号 成功返回结果
type TaobaotradessoldqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"trades_sold_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
	TidList []string `json:"tid_list,omitempty" xml:"tid_list>string,omitempty"`
}