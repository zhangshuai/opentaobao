package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradesSimpleSoldIncrementGetAPIResponse 查询卖家已卖出的增量交易简易数据（根据修改时间） API返回值
// taobao.trades.simple.sold.increment.get
//
// 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
// &lt;br/&gt;1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified &lt;= 1天。
// &lt;br/&gt;2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
// &lt;br/&gt;3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.simple.get获取订单详情。
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoTradesSimpleSoldIncrementGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradesSimpleSoldIncrementGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradesSimpleSoldIncrementGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradesSimpleSoldIncrementGetAPIResponseModel).Reset()
}

// TaobaoTradesSimpleSoldIncrementGetAPIResponseModel is 查询卖家已卖出的增量交易简易数据（根据修改时间） 成功返回结果
type TaobaoTradesSimpleSoldIncrementGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trades_simple_sold_increment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trades []Trade `json:"trades,omitempty" xml:"trades>trade,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradesSimpleSoldIncrementGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trades = m.Trades[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoTradesSimpleSoldIncrementGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradesSimpleSoldIncrementGetAPIResponse)
	},
}

// GetTaobaoTradesSimpleSoldIncrementGetAPIResponse 从 sync.Pool 获取 TaobaoTradesSimpleSoldIncrementGetAPIResponse
func GetTaobaoTradesSimpleSoldIncrementGetAPIResponse() *TaobaoTradesSimpleSoldIncrementGetAPIResponse {
	return poolTaobaoTradesSimpleSoldIncrementGetAPIResponse.Get().(*TaobaoTradesSimpleSoldIncrementGetAPIResponse)
}

// ReleaseTaobaoTradesSimpleSoldIncrementGetAPIResponse 将 TaobaoTradesSimpleSoldIncrementGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradesSimpleSoldIncrementGetAPIResponse(v *TaobaoTradesSimpleSoldIncrementGetAPIResponse) {
	v.Reset()
	poolTaobaoTradesSimpleSoldIncrementGetAPIResponse.Put(v)
}
