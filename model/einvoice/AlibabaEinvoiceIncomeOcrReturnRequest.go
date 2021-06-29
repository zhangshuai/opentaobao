package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传发票ocr的结果 API请求
alibaba.einvoice.income.ocr.return

服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
*/
type AlibabaEinvoiceIncomeOcrReturnRequest struct {
    model.Params
    // 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
    checksum   string
    // 错误码，success=false是必填
    errorCode   string
    // 错误消息，success=false是必填
    errorMessage   string
    // 发票ocr影像文件，type=1时必填
    imageData   []*model.File
    // 发票ocr影像编号，type=1时必填
    imageId   string
    // 发票代码，success=true时必填
    invoiceCode   string
    // 开票日期，格式为yyyy-MM-dd，success=true时必填
    invoiceDate   string
    // 发票种类，1=普票，2=专票，success=true时必填
    invoiceKind   int64
    // 发票号码，success=true时必填
    invoiceNo   string
    // 开票请求标识，扫描驱动回传type=1时填批次号
    reqIndex   string
    // ocr结果，true=成功，false=失败
    success   bool
    // 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
    sumPrice   string
    // 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
    type   int64
}

// 初始化AlibabaEinvoiceIncomeOcrReturnRequest对象
func NewAlibabaEinvoiceIncomeOcrReturnRequest() *AlibabaEinvoiceIncomeOcrReturnRequest{
    return &AlibabaEinvoiceIncomeOcrReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.ocr.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Checksum Setter
// 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetChecksum(checksum string) error {
    r.checksum = checksum
    r.Set("checksum", checksum)
    return nil
}

// Checksum Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetChecksum() string {
    return r.checksum
}
// ErrorCode Setter
// 错误码，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorCode() string {
    return r.errorCode
}
// ErrorMessage Setter
// 错误消息，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}
// ImageData Setter
// 发票ocr影像文件，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageData(imageData []*model.File) error {
    r.imageData = imageData
    r.Set("image_data", imageData)
    return nil
}

// ImageData Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageData() []*model.File {
    return r.imageData
}
// ImageId Setter
// 发票ocr影像编号，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageId(imageId string) error {
    r.imageId = imageId
    r.Set("image_id", imageId)
    return nil
}

// ImageId Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageId() string {
    return r.imageId
}
// InvoiceCode Setter
// 发票代码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// InvoiceDate Setter
// 开票日期，格式为yyyy-MM-dd，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}
// InvoiceKind Setter
// 发票种类，1=普票，2=专票，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}
// InvoiceNo Setter
// 发票号码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
// ReqIndex Setter
// 开票请求标识，扫描驱动回传type=1时填批次号
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetReqIndex() string {
    return r.reqIndex
}
// Success Setter
// ocr结果，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSuccess() bool {
    return r.success
}
// SumPrice Setter
// 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSumPrice() string {
    return r.sumPrice
}
// Type Setter
// 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetType() int64 {
    return r.type
}