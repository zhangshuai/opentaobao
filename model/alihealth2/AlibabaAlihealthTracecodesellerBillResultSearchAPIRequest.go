package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerbillresultsearchAPIRequest 查询出入库单处理结果api API请求
// alibaba.alihealth.tracecodeseller.bill.result.search
//
// 查询出入库单处理结果api
type AlibabaalihealthtracecodesellerbillresultsearchAPIRequest struct {
	model.Params
	// top身份认证
	_skeyCode string
	// 单据编号
	_billCode string
	// 查询开始日期
	_beginDate string
	// 查询结束日期
	_endDate string
	// 不需要
	_sellerName string
	// 商家id
	_entInfoId int64
	// 当前页
	_page int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaalihealthtracecodesellerbillresultsearchRequest 初始化AlibabaalihealthtracecodesellerbillresultsearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerbillresultsearchRequest() *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest {
	return &AlibabaalihealthtracecodesellerbillresultsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.result.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// top身份认证
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBeginDate is BeginDate Setter
// 查询开始日期
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 查询结束日期
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetSellerName is SellerName Setter
// 不需要
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetSellerName(_sellerName string) error {
	r._sellerName = _sellerName
	r.Set("seller_name", _sellerName)
	return nil
}

// GetSellerName SellerName Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetSellerName() string {
	return r._sellerName
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetPage is Page Setter
// 当前页
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthtracecodesellerbillresultsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
