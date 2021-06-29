package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易可购买用户标记 API返回值 
taobao.opentrade.queue.users.mark

尖货交易用户标记信息回传，根据openId标记用户可购买商品
*/
type TaobaoOpentradeQueueUsersMarkAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeQueueUsersMarkResponse
}

// 尖货交易可购买用户标记 成功返回结果
type TaobaoOpentradeQueueUsersMarkResponse struct {
    XMLName xml.Name `xml:"opentrade_queue_users_mark_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 标记成功的用户数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
