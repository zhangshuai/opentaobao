package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailcommissionorderqueryAPIRequest 分销订单查询 API请求
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
type AlibabaretailcommissionorderqueryAPIRequest struct {
	model.Params
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	_endPayTime string
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	_startPayTime string
	// 页码，默认第一页
	_pageNo int64
	// 页大小，默认每页十条
	_pageSize int64
}

// NewAlibabaretailcommissionorderqueryRequest 初始化AlibabaretailcommissionorderqueryAPIRequest对象
func NewAlibabaretailcommissionorderqueryRequest() *AlibabaretailcommissionorderqueryAPIRequest {
	return &AlibabaretailcommissionorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailcommissionorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailcommissionorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailcommissionorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndPayTime is EndPayTime Setter
// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaretailcommissionorderqueryAPIRequest) SetEndPayTime(_endPayTime string) error {
	r._endPayTime = _endPayTime
	r.Set("end_pay_time", _endPayTime)
	return nil
}

// GetEndPayTime EndPayTime Getter
func (r AlibabaretailcommissionorderqueryAPIRequest) GetEndPayTime() string {
	return r._endPayTime
}

// SetStartPayTime is StartPayTime Setter
// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaretailcommissionorderqueryAPIRequest) SetStartPayTime(_startPayTime string) error {
	r._startPayTime = _startPayTime
	r.Set("start_pay_time", _startPayTime)
	return nil
}

// GetStartPayTime StartPayTime Getter
func (r AlibabaretailcommissionorderqueryAPIRequest) GetStartPayTime() string {
	return r._startPayTime
}

// SetPageNo is PageNo Setter
// 页码，默认第一页
func (r *AlibabaretailcommissionorderqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaretailcommissionorderqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小，默认每页十条
func (r *AlibabaretailcommissionorderqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaretailcommissionorderqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
