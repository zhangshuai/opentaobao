package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票结算单同步前开发票 API请求
alibaba.einvoice.bill.forword.create

扫码开票结算单同步前开发票，会将数据同步到结算单中
*/
type AlibabaEinvoiceBillForwordCreateRequest struct {
    model.Params
    // 发票商品明细
    invoiceItems   []BillItemDO
    // 结算单订单ID
    orderId   string
    // 收款方税号
    payeeRegisterNo   string
    // 调用平台，用以区分不同的订单ID，不填默认为default
    platform   string
    // 开票流水号，若不填则系统默认生成
    seriNo   string
    // 店铺名称
    shopName   string
    // 开票金额，和明细累计总金额需相同
    sumPrice   string
    // 付款方地址
    payerAddress   string
    // 付款方银行账户
    payerBankaccount   string
    // 付款方开票邮件通知邮箱
    payerEmail   string
    // 付款方发票抬头
    payerName   string
    // 企业电话
    payerPhone   string
    // 付款方税号
    payerRegisterNo   string
    // 个人电话，接收发票通知
    phoneNumber   string
    // 企业或个人抬头发票，0=个人，1=企业。默认=1
    businessType   int64
}

// 初始化AlibabaEinvoiceBillForwordCreateRequest对象
func NewAlibabaEinvoiceBillForwordCreateRequest() *AlibabaEinvoiceBillForwordCreateRequest{
    return &AlibabaEinvoiceBillForwordCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceBillForwordCreateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.bill.forword.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceBillForwordCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceItems Setter
// 发票商品明细
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetInvoiceItems(invoiceItems []BillItemDO) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetInvoiceItems() []BillItemDO {
    return r.invoiceItems
}
// OrderId Setter
// 结算单订单ID
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetOrderId() string {
    return r.orderId
}
// PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// Platform Setter
// 调用平台，用以区分不同的订单ID，不填默认为default
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPlatform() string {
    return r.platform
}
// SeriNo Setter
// 开票流水号，若不填则系统默认生成
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetSeriNo(seriNo string) error {
    r.seriNo = seriNo
    r.Set("seri_no", seriNo)
    return nil
}

// SeriNo Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetSeriNo() string {
    return r.seriNo
}
// ShopName Setter
// 店铺名称
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetShopName(shopName string) error {
    r.shopName = shopName
    r.Set("shop_name", shopName)
    return nil
}

// ShopName Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetShopName() string {
    return r.shopName
}
// SumPrice Setter
// 开票金额，和明细累计总金额需相同
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetSumPrice() string {
    return r.sumPrice
}
// PayerAddress Setter
// 付款方地址
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerAddress(payerAddress string) error {
    r.payerAddress = payerAddress
    r.Set("payer_address", payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerAddress() string {
    return r.payerAddress
}
// PayerBankaccount Setter
// 付款方银行账户
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerBankaccount(payerBankaccount string) error {
    r.payerBankaccount = payerBankaccount
    r.Set("payer_bankaccount", payerBankaccount)
    return nil
}

// PayerBankaccount Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerBankaccount() string {
    return r.payerBankaccount
}
// PayerEmail Setter
// 付款方开票邮件通知邮箱
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerEmail(payerEmail string) error {
    r.payerEmail = payerEmail
    r.Set("payer_email", payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerEmail() string {
    return r.payerEmail
}
// PayerName Setter
// 付款方发票抬头
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerName() string {
    return r.payerName
}
// PayerPhone Setter
// 企业电话
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerPhone(payerPhone string) error {
    r.payerPhone = payerPhone
    r.Set("payer_phone", payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerPhone() string {
    return r.payerPhone
}
// PayerRegisterNo Setter
// 付款方税号
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}
// PhoneNumber Setter
// 个人电话，接收发票通知
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetPhoneNumber() string {
    return r.phoneNumber
}
// BusinessType Setter
// 企业或个人抬头发票，0=个人，1=企业。默认=1
func (r *AlibabaEinvoiceBillForwordCreateRequest) SetBusinessType(businessType int64) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceBillForwordCreateRequest) GetBusinessType() int64 {
    return r.businessType
}