package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔交易的详细信息 APIResponse
taobao.trade.fullinfo.get

获取单笔交易的详细信息
<br/>1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
<br/>2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>3. 获取红包优惠金额可以使用字段 coupon_fee
<br/>4. 请按需获取字段，减少TOP系统的压力
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
type TaobaoTradeFullinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeFullinfoGetResponse `json:"taobao_trade_fullinfo_get_response,omitempty"`
}

type TaobaoTradeFullinfoGetResponse struct {

    // 交易主订单信息
    Trade   *Trade `json:"trade,omitempty"`

}