package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignScheduleGetAPIRequest 取得一个推广计划的分时折扣设置 API请求
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
type TaobaoSimbaCampaignScheduleGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignScheduleGetRequest 初始化TaobaoSimbaCampaignScheduleGetAPIRequest对象
func NewTaobaoSimbaCampaignScheduleGetRequest() *TaobaoSimbaCampaignScheduleGetAPIRequest {
	return &TaobaoSimbaCampaignScheduleGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignScheduleGetAPIRequest) Reset() {
	r._nick = ""
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignScheduleGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignScheduleGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignScheduleGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaCampaignScheduleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignScheduleGetRequest()
	},
}

// GetTaobaoSimbaCampaignScheduleGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignScheduleGetAPIRequest
func GetTaobaoSimbaCampaignScheduleGetAPIRequest() *TaobaoSimbaCampaignScheduleGetAPIRequest {
	return poolTaobaoSimbaCampaignScheduleGetAPIRequest.Get().(*TaobaoSimbaCampaignScheduleGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignScheduleGetAPIRequest 将 TaobaoSimbaCampaignScheduleGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignScheduleGetAPIRequest(v *TaobaoSimbaCampaignScheduleGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignScheduleGetAPIRequest.Put(v)
}
