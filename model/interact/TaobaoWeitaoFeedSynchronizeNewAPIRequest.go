package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed API请求
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘
*/
type TaobaoWeitaoFeedSynchronizeNewAPIRequest struct {
    model.Params
    // 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
    _feedType   int64
    // feed点击跳转的活动地址
    _detailUrl   string
    // feed展示结束时间
    _endTime   int64
    // feed展示开始时间
    _startTime   int64
    // feed描述
    _summary   string
    // feed标题
    _title   string
    // 宝贝列表，用于card展示，0~2个宝贝ID
    _itemIds   []string
    // 活动ID
    _sbizId   string
}

// 初始化TaobaoWeitaoFeedSynchronizeNewAPIRequest对象
func NewTaobaoWeitaoFeedSynchronizeNewRequest() *TaobaoWeitaoFeedSynchronizeNewAPIRequest{
    return &TaobaoWeitaoFeedSynchronizeNewAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.synchronize.new"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FeedType Setter
// 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetFeedType(_feedType int64) error {
    r._feedType = _feedType
    r.Set("feed_type", _feedType)
    return nil
}

// FeedType Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetFeedType() int64 {
    return r._feedType
}
// DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetDetailUrl(_detailUrl string) error {
    r._detailUrl = _detailUrl
    r.Set("detail_url", _detailUrl)
    return nil
}

// DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetDetailUrl() string {
    return r._detailUrl
}
// EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetEndTime(_endTime int64) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetEndTime() int64 {
    return r._endTime
}
// StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetStartTime(_startTime int64) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetStartTime() int64 {
    return r._startTime
}
// Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetSummary(_summary string) error {
    r._summary = _summary
    r.Set("summary", _summary)
    return nil
}

// Summary Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetSummary() string {
    return r._summary
}
// Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetTitle() string {
    return r._title
}
// ItemIds Setter
// 宝贝列表，用于card展示，0~2个宝贝ID
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetItemIds(_itemIds []string) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetItemIds() []string {
    return r._itemIds
}
// SbizId Setter
// 活动ID
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetSbizId(_sbizId string) error {
    r._sbizId = _sbizId
    r.Set("sbiz_id", _sbizId)
    return nil
}

// SbizId Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetSbizId() string {
    return r._sbizId
}