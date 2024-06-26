package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderGetAPIResponse 代理商获取订单信息回调API API返回值
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderGetAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderGetAPIResponseModel is 代理商获取订单信息回调API 成功返回结果
type TaobaoTrainAgentOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 火车票信息。
	Tickets []ToAgentTicketInfo `json:"tickets,omitempty" xml:"tickets>to_agent_ticket_info,omitempty"`
	// 移动电话
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 保险邮件地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 联系人姓名
	RelationName string `json:"relation_name,omitempty" xml:"relation_name,omitempty"`
	// 如果是公司发票，需要公司名称，如果不需要公司名称，返回no
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 最晚出票时间
	LatestIssueTime string `json:"latest_issue_time,omitempty" xml:"latest_issue_time,omitempty"`
	// 线下票收件人姓名
	TransportName string `json:"transport_name,omitempty" xml:"transport_name,omitempty"`
	// 线下票收件人手机号
	TransportPhone string `json:"transport_phone,omitempty" xml:"transport_phone,omitempty"`
	// 线下票收件人地址
	TransportAddress string `json:"transport_address,omitempty" xml:"transport_address,omitempty"`
	// 扩展字段
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// ttp_order_id
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 整个订单的总价,包括每张票价及保险价格,价格精确到分,例如100元,输出为10000.
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 1-已付款，2-关闭，3-成功
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单类型0:默认订单类型走代理商账号；1：走12306客户绑定的账号；2：线下邮寄票
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 纸质票类型: 1 靠窗,2 连坐,3 上铺,4 中铺,5 下铺,6 是否同包厢
	PaperType int64 `json:"paper_type,omitempty" xml:"paper_type,omitempty"`
	// 当下铺/靠窗/连坐无票时,是否支持非下铺/非靠窗/非连坐(0不接受,1接受)
	PaperBackup int64 `json:"paper_backup,omitempty" xml:"paper_backup,omitempty"`
	// 至少接受下铺/靠窗/连坐数量
	PaperLowSeatCount int64 `json:"paper_low_seat_count,omitempty" xml:"paper_low_seat_count,omitempty"`
	// 快递费（分）
	TransportPrice int64 `json:"transport_price,omitempty" xml:"transport_price,omitempty"`
	// 手续费总价（分）
	ServicePrice int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// 返回错误。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否需要保险邮件地址
	Mailing bool `json:"mailing,omitempty" xml:"mailing,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tickets = m.Tickets[:0]
	m.Telephone = ""
	m.Address = ""
	m.RelationName = ""
	m.CompanyName = ""
	m.LatestIssueTime = ""
	m.TransportName = ""
	m.TransportPhone = ""
	m.TransportAddress = ""
	m.ExtendParams = ""
	m.TtpOrderId = 0
	m.MainOrderId = 0
	m.TotalPrice = 0
	m.OrderStatus = 0
	m.OrderType = 0
	m.PaperType = 0
	m.PaperBackup = 0
	m.PaperLowSeatCount = 0
	m.TransportPrice = 0
	m.ServicePrice = 0
	m.IsSuccess = false
	m.Mailing = false
}

var poolTaobaoTrainAgentOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderGetAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderGetAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderGetAPIResponse
func GetTaobaoTrainAgentOrderGetAPIResponse() *TaobaoTrainAgentOrderGetAPIResponse {
	return poolTaobaoTrainAgentOrderGetAPIResponse.Get().(*TaobaoTrainAgentOrderGetAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderGetAPIResponse 将 TaobaoTrainAgentOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderGetAPIResponse(v *TaobaoTrainAgentOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderGetAPIResponse.Put(v)
}
