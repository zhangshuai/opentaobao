package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 API请求
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetRequest struct {
    model.Params
    // 权限校验参数
    subwayToken   string
    // 昵称
    nick   string
    // 推广计划id
    campaignId   int64
    // 推广组id
    adgroupId   int64
    // 开始时间，格式yyyy-mm-dd
    startTime   string
    // 结束时间，格式yyyy-mm-dd
    endTime   string
    // 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
    searchType   string
    // 页码
    pageNo   int64
    // 每页大小
    pageSize   int64
    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
    source   string
}

// 初始化TaobaoSimbaRptAdgroupbaseGetRequest对象
func NewTaobaoSimbaRptAdgroupbaseGetRequest() *TaobaoSimbaRptAdgroupbaseGetRequest{
    return &TaobaoSimbaRptAdgroupbaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetEndTime() string {
    return r.endTime
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSearchType() string {
    return r.searchType
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSource() string {
    return r.source
}
