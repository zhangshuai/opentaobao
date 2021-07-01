package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单列表 API请求
tmall.channel.trade.deliverorder.gets

查询发货单列表
*/
type TmallChannelTradeDeliverorderGetsAPIRequest struct {
    model.Params
    // 发货单单号
    _mainDeliverOrderNo   int64
    // 发货单状态列表
    _orderStatusList   []int64
    // 是否包括子发货单
    _isIncludeSubOrder   bool
    // 每页显示数量
    _pageSize   int64
    // 查询第几页
    _pageNumber   int64
    // 是否分页查询
    _needPagination   bool
    // 渠道
    _channel   int64
}

// 初始化TmallChannelTradeDeliverorderGetsAPIRequest对象
func NewTmallChannelTradeDeliverorderGetsRequest() *TmallChannelTradeDeliverorderGetsAPIRequest{
    return &TmallChannelTradeDeliverorderGetsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.gets"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单单号
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// OrderStatusList Setter
// 发货单状态列表
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetOrderStatusList(_orderStatusList []int64) error {
    r._orderStatusList = _orderStatusList
    r.Set("order_status_list", _orderStatusList)
    return nil
}

// OrderStatusList Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetOrderStatusList() []int64 {
    return r._orderStatusList
}
// IsIncludeSubOrder Setter
// 是否包括子发货单
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
// PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetPageNumber() int64 {
    return r._pageNumber
}
// NeedPagination Setter
// 是否分页查询
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetNeedPagination(_needPagination bool) error {
    r._needPagination = _needPagination
    r.Set("need_pagination", _needPagination)
    return nil
}

// NeedPagination Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetNeedPagination() bool {
    return r._needPagination
}
// Channel Setter
// 渠道
func (r *TmallChannelTradeDeliverorderGetsAPIRequest) SetChannel(_channel int64) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TmallChannelTradeDeliverorderGetsAPIRequest) GetChannel() int64 {
    return r._channel
}