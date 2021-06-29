package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划分日数据查询 API请求
taobao.feedflow.item.campaign.rptdailylist

推广计划分日数据查询
*/
type TaobaoFeedflowItemCampaignRptdailylistRequest struct {
    model.Params
    // 查询条件
    rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCampaignRptdailylistRequest对象
func NewTaobaoFeedflowItemCampaignRptdailylistRequest() *TaobaoFeedflowItemCampaignRptdailylistRequest{
    return &TaobaoFeedflowItemCampaignRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCampaignRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}