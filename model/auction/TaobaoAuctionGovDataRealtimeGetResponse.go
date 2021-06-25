package auction

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取实时(今日)统计数据 APIResponse
taobao.auction.gov.data.realtime.get

提供查询当日法院及下属法院的拍卖统计数据
*/
type TaobaoAuctionGovDataRealtimeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAuctionGovDataRealtimeGetResponse `json:"taobao_auction_gov_data_realtime_get_response,omitempty"`
}

type TaobaoAuctionGovDataRealtimeGetResponse struct {

    // 今日拍卖统计数据
    RealTimeData   *RealTimeData `json:"real_time_data,omitempty"`

}