package alscmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketcodeUseAPIResponse 外部核销服务 API返回值
// alibaba.alsc.merchant.ext.ticketcode.use
//
// 外部核销服务
type AlibabaAlscMerchantExtTicketcodeUseAPIResponse struct {
	model.CommonResponse
	AlibabaAlscMerchantExtTicketcodeUseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscMerchantExtTicketcodeUseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscMerchantExtTicketcodeUseAPIResponseModel).Reset()
}

// AlibabaAlscMerchantExtTicketcodeUseAPIResponseModel is 外部核销服务 成功返回结果
type AlibabaAlscMerchantExtTicketcodeUseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_merchant_ext_ticketcode_use_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一次性核销多份场景，被核销的凭证明细信息
	TicketUseDetails []MerchantTicketUseDetail `json:"ticket_use_details,omitempty" xml:"ticket_use_details>merchant_ticket_use_detail,omitempty"`
	// 口碑子订单号, 即alsc_sub_order_no /***** 20220922后对接的服务商，不支持该字段 ******
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 用户购买时商品的原价，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 该字段用于描述本次返回中的业务属性，现有：TICKET_ALREADY_USED（幂等业务码）
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 用户购买时商品的现价，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	CurrentPrice string `json:"current_price,omitempty" xml:"current_price,omitempty"`
	// 商家优惠金额，单位为元，精确到小数点后两位，一次性核销多份券场景，返回总商家优惠金额 /***** 20220922后对接的服务商，不支持该字段 ******
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 交易中可给用户开具发票的金额，单位为元，精确到小数点后两位，一次性核销多份券场景，返回总开票金额 /***** 20220922后对接的服务商，不支持该字段 ******
	InvoiceAmount string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 券核销门店ID /***** 20220922后对接的服务商，不支持该字段 ******
	UseShopId string `json:"use_shop_id,omitempty" xml:"use_shop_id,omitempty"`
	// 凭证对应商品别名，在口碑商品创建时候设置 /***** 20220922后对接的服务商，不支持该字段 ******
	ItemAlias string `json:"item_alias,omitempty" xml:"item_alias,omitempty"`
	// 券核销门店名称 /***** 20220922后对接的服务商，不支持该字段 ******
	UseShopName string `json:"use_shop_name,omitempty" xml:"use_shop_name,omitempty"`
	// 口碑商品ID /***** 20220922后对接的服务商，不支持该字段 ******
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家实收金额，单位为元，精确到小数点后两位，一次性核销多份券场景，返回总商家实收金额 /***** 20220922后对接的服务商，不支持该字段 ******
	ReceiptAmount string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 用户购买时商品的名称 /***** 20220922后对接的服务商，不支持该字段 ******
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 核销流水号 /***** 20220922后对接的服务商，不支持该字段 ******
	TicketTransId string `json:"ticket_trans_id,omitempty" xml:"ticket_trans_id,omitempty"`
	// 券核销时间
	UseDate string `json:"use_date,omitempty" xml:"use_date,omitempty"`
	// 凭证码对应的凭证资产id /***** 20220922后对接的服务商，不支持该字段 ******
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
	TicketRequestId string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`
	// 用户购买券的时候实际支付的金额，单位为元，精确到小数点后两位，一次性核销多份券场景，返回总实际支付金额 /***** 20220922后对接的服务商，不支持该字段 ******
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" xml:"buyer_pay_amount,omitempty"`
	// 口碑补贴金额，单位为元，精确到小数点后两位，一次性核销多份券场景，返回总口碑补贴金额 /***** 20220922后对接的服务商，不支持该字段 ******
	MerchantSubsidyAmount string `json:"merchant_subsidy_amount,omitempty" xml:"merchant_subsidy_amount,omitempty"`
	// 12位的券码，券码为纯数字，且唯一不重复 /***** 20220922后对接的服务商，不支持该字段 ******
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscMerchantExtTicketcodeUseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TicketUseDetails = m.TicketUseDetails[:0]
	m.OrderNo = ""
	m.OriginalPrice = ""
	m.BizCode = ""
	m.CurrentPrice = ""
	m.DiscountAmount = ""
	m.InvoiceAmount = ""
	m.UseShopId = ""
	m.ItemAlias = ""
	m.UseShopName = ""
	m.ItemId = ""
	m.ReceiptAmount = ""
	m.ItemName = ""
	m.TicketTransId = ""
	m.UseDate = ""
	m.VoucherId = ""
	m.TicketRequestId = ""
	m.BuyerPayAmount = ""
	m.MerchantSubsidyAmount = ""
	m.TicketCode = ""
}

var poolAlibabaAlscMerchantExtTicketcodeUseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscMerchantExtTicketcodeUseAPIResponse)
	},
}

// GetAlibabaAlscMerchantExtTicketcodeUseAPIResponse 从 sync.Pool 获取 AlibabaAlscMerchantExtTicketcodeUseAPIResponse
func GetAlibabaAlscMerchantExtTicketcodeUseAPIResponse() *AlibabaAlscMerchantExtTicketcodeUseAPIResponse {
	return poolAlibabaAlscMerchantExtTicketcodeUseAPIResponse.Get().(*AlibabaAlscMerchantExtTicketcodeUseAPIResponse)
}

// ReleaseAlibabaAlscMerchantExtTicketcodeUseAPIResponse 将 AlibabaAlscMerchantExtTicketcodeUseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscMerchantExtTicketcodeUseAPIResponse(v *AlibabaAlscMerchantExtTicketcodeUseAPIResponse) {
	v.Reset()
	poolAlibabaAlscMerchantExtTicketcodeUseAPIResponse.Put(v)
}
