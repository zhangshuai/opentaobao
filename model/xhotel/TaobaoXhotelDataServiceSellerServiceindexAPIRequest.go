package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDataServiceSellerServiceindexAPIRequest 卖家服务指数查询 API请求
// taobao.xhotel.data.service.seller.serviceindex
//
// 卖家服务指数查询
type TaobaoXhotelDataServiceSellerServiceindexAPIRequest struct {
	model.Params
	// 渠道商名称
	_vendor string
	// 分页参数
	_startRow int64
	// 分页参数
	_pageSize int64
	// 查询截止日期
	_reportEndDate string
	// 查询开始日期
	_reportStartDate string
}

// NewTaobaoXhotelDataServiceSellerServiceindexRequest 初始化TaobaoXhotelDataServiceSellerServiceindexAPIRequest对象
func NewTaobaoXhotelDataServiceSellerServiceindexRequest() *TaobaoXhotelDataServiceSellerServiceindexAPIRequest {
	return &TaobaoXhotelDataServiceSellerServiceindexAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.data.service.seller.serviceindex"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceSellerServiceindexAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is StartRow Setter
// 分页参数
func (r *TaobaoXhotelDataServiceSellerServiceindexAPIRequest) SetStartRow(_startRow int64) error {
	r._startRow = _startRow
	r.Set("start_row", _startRow)
	return nil
}

// Get StartRow Getter
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetStartRow() int64 {
	return r._startRow
}

// Set is PageSize Setter
// 分页参数
func (r *TaobaoXhotelDataServiceSellerServiceindexAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ReportEndDate Setter
// 查询截止日期
func (r *TaobaoXhotelDataServiceSellerServiceindexAPIRequest) SetReportEndDate(_reportEndDate string) error {
	r._reportEndDate = _reportEndDate
	r.Set("report_end_date", _reportEndDate)
	return nil
}

// Get ReportEndDate Getter
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetReportEndDate() string {
	return r._reportEndDate
}

// Set is ReportStartDate Setter
// 查询开始日期
func (r *TaobaoXhotelDataServiceSellerServiceindexAPIRequest) SetReportStartDate(_reportStartDate string) error {
	r._reportStartDate = _reportStartDate
	r.Set("report_start_date", _reportStartDate)
	return nil
}

// Get ReportStartDate Getter
func (r TaobaoXhotelDataServiceSellerServiceindexAPIRequest) GetReportStartDate() string {
	return r._reportStartDate
}
