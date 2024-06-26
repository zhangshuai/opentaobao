package tmallchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeOrderGetsAPIRequest 分页查询采购单 API请求
// tmall.channel.trade.order.gets
//
// 分页查询采购单
type TmallChannelTradeOrderGetsAPIRequest struct {
	model.Params
	// 分销商Nick
	_distributorNick string
	// 创建时间从
	_createTimeStart string
	// 创建时间到
	_createTimeEnd string
	// 每页显示数量
	_pageSize int64
	// 查询第几页
	_pageNumber int64
	// 主采购单号
	_mainPurchaseOrderNo int64
	// 渠道编码
	_channel int64
	// 1-代销；2-经销
	_tradeType int64
	// 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
	_orderStatus int64
	// 是否包含子单
	_isIncludeSubOrder bool
	// 是否包含主单
	_isIncludeMainOrder bool
	// 是否包含物流信息
	_isIncludeLogistics bool
	// 是否分页查询
	_needPagination bool
}

// NewTmallChannelTradeOrderGetsRequest 初始化TmallChannelTradeOrderGetsAPIRequest对象
func NewTmallChannelTradeOrderGetsRequest() *TmallChannelTradeOrderGetsAPIRequest {
	return &TmallChannelTradeOrderGetsAPIRequest{
		Params: model.NewParams(13),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallChannelTradeOrderGetsAPIRequest) Reset() {
	r._distributorNick = ""
	r._createTimeStart = ""
	r._createTimeEnd = ""
	r._pageSize = 0
	r._pageNumber = 0
	r._mainPurchaseOrderNo = 0
	r._channel = 0
	r._tradeType = 0
	r._orderStatus = 0
	r._isIncludeSubOrder = false
	r._isIncludeMainOrder = false
	r._isIncludeLogistics = false
	r._needPagination = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderGetsAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.order.gets"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderGetsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallChannelTradeOrderGetsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorNick is DistributorNick Setter
// 分销商Nick
func (r *TmallChannelTradeOrderGetsAPIRequest) SetDistributorNick(_distributorNick string) error {
	r._distributorNick = _distributorNick
	r.Set("distributor_nick", _distributorNick)
	return nil
}

// GetDistributorNick DistributorNick Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetDistributorNick() string {
	return r._distributorNick
}

// SetCreateTimeStart is CreateTimeStart Setter
// 创建时间从
func (r *TmallChannelTradeOrderGetsAPIRequest) SetCreateTimeStart(_createTimeStart string) error {
	r._createTimeStart = _createTimeStart
	r.Set("create_time_start", _createTimeStart)
	return nil
}

// GetCreateTimeStart CreateTimeStart Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetCreateTimeStart() string {
	return r._createTimeStart
}

// SetCreateTimeEnd is CreateTimeEnd Setter
// 创建时间到
func (r *TmallChannelTradeOrderGetsAPIRequest) SetCreateTimeEnd(_createTimeEnd string) error {
	r._createTimeEnd = _createTimeEnd
	r.Set("create_time_end", _createTimeEnd)
	return nil
}

// GetCreateTimeEnd CreateTimeEnd Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetCreateTimeEnd() string {
	return r._createTimeEnd
}

// SetPageSize is PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeOrderGetsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeOrderGetsAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

// SetMainPurchaseOrderNo is MainPurchaseOrderNo Setter
// 主采购单号
func (r *TmallChannelTradeOrderGetsAPIRequest) SetMainPurchaseOrderNo(_mainPurchaseOrderNo int64) error {
	r._mainPurchaseOrderNo = _mainPurchaseOrderNo
	r.Set("main_purchase_order_no", _mainPurchaseOrderNo)
	return nil
}

// GetMainPurchaseOrderNo MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetMainPurchaseOrderNo() int64 {
	return r._mainPurchaseOrderNo
}

// SetChannel is Channel Setter
// 渠道编码
func (r *TmallChannelTradeOrderGetsAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetChannel() int64 {
	return r._channel
}

// SetTradeType is TradeType Setter
// 1-代销；2-经销
func (r *TmallChannelTradeOrderGetsAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetTradeType() int64 {
	return r._tradeType
}

// SetOrderStatus is OrderStatus Setter
// 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
func (r *TmallChannelTradeOrderGetsAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetIsIncludeSubOrder is IsIncludeSubOrder Setter
// 是否包含子单
func (r *TmallChannelTradeOrderGetsAPIRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
	r._isIncludeSubOrder = _isIncludeSubOrder
	r.Set("is_include_sub_order", _isIncludeSubOrder)
	return nil
}

// GetIsIncludeSubOrder IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetIsIncludeSubOrder() bool {
	return r._isIncludeSubOrder
}

// SetIsIncludeMainOrder is IsIncludeMainOrder Setter
// 是否包含主单
func (r *TmallChannelTradeOrderGetsAPIRequest) SetIsIncludeMainOrder(_isIncludeMainOrder bool) error {
	r._isIncludeMainOrder = _isIncludeMainOrder
	r.Set("is_include_main_order", _isIncludeMainOrder)
	return nil
}

// GetIsIncludeMainOrder IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetIsIncludeMainOrder() bool {
	return r._isIncludeMainOrder
}

// SetIsIncludeLogistics is IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetsAPIRequest) SetIsIncludeLogistics(_isIncludeLogistics bool) error {
	r._isIncludeLogistics = _isIncludeLogistics
	r.Set("is_include_logistics", _isIncludeLogistics)
	return nil
}

// GetIsIncludeLogistics IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetIsIncludeLogistics() bool {
	return r._isIncludeLogistics
}

// SetNeedPagination is NeedPagination Setter
// 是否分页查询
func (r *TmallChannelTradeOrderGetsAPIRequest) SetNeedPagination(_needPagination bool) error {
	r._needPagination = _needPagination
	r.Set("need_pagination", _needPagination)
	return nil
}

// GetNeedPagination NeedPagination Getter
func (r TmallChannelTradeOrderGetsAPIRequest) GetNeedPagination() bool {
	return r._needPagination
}

var poolTmallChannelTradeOrderGetsAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallChannelTradeOrderGetsRequest()
	},
}

// GetTmallChannelTradeOrderGetsRequest 从 sync.Pool 获取 TmallChannelTradeOrderGetsAPIRequest
func GetTmallChannelTradeOrderGetsAPIRequest() *TmallChannelTradeOrderGetsAPIRequest {
	return poolTmallChannelTradeOrderGetsAPIRequest.Get().(*TmallChannelTradeOrderGetsAPIRequest)
}

// ReleaseTmallChannelTradeOrderGetsAPIRequest 将 TmallChannelTradeOrderGetsAPIRequest 放入 sync.Pool
func ReleaseTmallChannelTradeOrderGetsAPIRequest(v *TmallChannelTradeOrderGetsAPIRequest) {
	v.Reset()
	poolTmallChannelTradeOrderGetsAPIRequest.Put(v)
}
