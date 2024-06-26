package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCreateResultsIncrementGetAPIRequest ERP增量开票结果获取 API请求
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetAPIRequest struct {
	model.Params
	// 终止查询时间
	_endModified string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 起始查询时间
	_startModified string
	// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
	_status string
	// 页面大小(不能超过200)
	_pageSize int64
	// 显示的页码
	_pageNo int64
}

// NewAlibabaEinvoiceCreateResultsIncrementGetRequest 初始化AlibabaEinvoiceCreateResultsIncrementGetAPIRequest对象
func NewAlibabaEinvoiceCreateResultsIncrementGetRequest() *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest {
	return &AlibabaEinvoiceCreateResultsIncrementGetAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) Reset() {
	r._endModified = ""
	r._payeeRegisterNo = ""
	r._startModified = ""
	r._status = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.create.results.increment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndModified is EndModified Setter
// 终止查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetStartModified is StartModified Setter
// 起始查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetStatus is Status Setter
// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetStatus() string {
	return r._status
}

// SetPageSize is PageSize Setter
// 页面大小(不能超过200)
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 显示的页码
func (r *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolAlibabaEinvoiceCreateResultsIncrementGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceCreateResultsIncrementGetRequest()
	},
}

// GetAlibabaEinvoiceCreateResultsIncrementGetRequest 从 sync.Pool 获取 AlibabaEinvoiceCreateResultsIncrementGetAPIRequest
func GetAlibabaEinvoiceCreateResultsIncrementGetAPIRequest() *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest {
	return poolAlibabaEinvoiceCreateResultsIncrementGetAPIRequest.Get().(*AlibabaEinvoiceCreateResultsIncrementGetAPIRequest)
}

// ReleaseAlibabaEinvoiceCreateResultsIncrementGetAPIRequest 将 AlibabaEinvoiceCreateResultsIncrementGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceCreateResultsIncrementGetAPIRequest(v *AlibabaEinvoiceCreateResultsIncrementGetAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceCreateResultsIncrementGetAPIRequest.Put(v)
}
