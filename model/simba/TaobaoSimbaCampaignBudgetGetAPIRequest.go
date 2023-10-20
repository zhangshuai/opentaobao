package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignbudgetgetAPIRequest 取得一个推广计划的日限额 API请求
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
type TaobaosimbacampaignbudgetgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignbudgetgetRequest 初始化TaobaosimbacampaignbudgetgetAPIRequest对象
func NewTaobaosimbacampaignbudgetgetRequest() *TaobaosimbacampaignbudgetgetAPIRequest {
	return &TaobaosimbacampaignbudgetgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignbudgetgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.budget.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignbudgetgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignbudgetgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignbudgetgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignbudgetgetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignbudgetgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignbudgetgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
