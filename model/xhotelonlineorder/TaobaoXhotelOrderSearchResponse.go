package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库订单查询 API返回值 
taobao.xhotel.order.search

酒店产品库订单查询功能，查询90天内的订单
*/
type TaobaoXhotelOrderSearchAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderSearchResponse
}

// 酒店产品库订单查询 成功返回结果
type TaobaoXhotelOrderSearchResponse struct {
    XMLName xml.Name `xml:"xhotel_order_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 订单信息
    HotelOrders   []XHotelOrder `json:"hotel_orders,omitempty" xml:"hotel_orders>x_hotel_order,omitempty"`
    // 符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
