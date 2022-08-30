package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljTemporaryCreateAPIRequest 淘宝客-推广者-淘礼金创建测试营销ID API请求
// taobao.tbk.dg.vegas.tlj.temporary.create
//
// 淘宝客-推广者-淘礼金创建测试营销ID
type TaobaoTbkDgVegasTljTemporaryCreateAPIRequest struct {
	model.Params
	// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	_useStartTime string
	// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	_useEndTime string
	// 发放截止时间
	_sendEndTime string
	// 发放开始时间
	_sendStartTime string
	// 单个淘礼金面额，支持两位小数，单位元
	_perFace string
	// 淘礼金名称，最大10个字符
	_name string
	// 宝贝ID或营销ID
	_itemId string
	// CPS佣金类型
	_campaignType string
	// 必须传入0
	_securityLevel int64
	// 结束日期的模式,1:相对时间，2:绝对时间
	_useEndTimeMode int64
	// 单用户累计中奖次数上限
	_userTotalWinNumLimit int64
	// 淘礼金总个数
	_totalNum int64
	// 妈妈广告位Id
	_adzoneId int64
	// 必须设置为true，默认开启安全
	_securitySwitch bool
}

// NewTaobaoTbkDgVegasTljTemporaryCreateRequest 初始化TaobaoTbkDgVegasTljTemporaryCreateAPIRequest对象
func NewTaobaoTbkDgVegasTljTemporaryCreateRequest() *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest {
	return &TaobaoTbkDgVegasTljTemporaryCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.temporary.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUseStartTime is UseStartTime Setter
// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetUseStartTime(_useStartTime string) error {
	r._useStartTime = _useStartTime
	r.Set("use_start_time", _useStartTime)
	return nil
}

// GetUseStartTime UseStartTime Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetUseStartTime() string {
	return r._useStartTime
}

// SetUseEndTime is UseEndTime Setter
// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetUseEndTime(_useEndTime string) error {
	r._useEndTime = _useEndTime
	r.Set("use_end_time", _useEndTime)
	return nil
}

// GetUseEndTime UseEndTime Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetUseEndTime() string {
	return r._useEndTime
}

// SetSendEndTime is SendEndTime Setter
// 发放截止时间
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetSendEndTime(_sendEndTime string) error {
	r._sendEndTime = _sendEndTime
	r.Set("send_end_time", _sendEndTime)
	return nil
}

// GetSendEndTime SendEndTime Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetSendEndTime() string {
	return r._sendEndTime
}

// SetSendStartTime is SendStartTime Setter
// 发放开始时间
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetSendStartTime(_sendStartTime string) error {
	r._sendStartTime = _sendStartTime
	r.Set("send_start_time", _sendStartTime)
	return nil
}

// GetSendStartTime SendStartTime Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetSendStartTime() string {
	return r._sendStartTime
}

// SetPerFace is PerFace Setter
// 单个淘礼金面额，支持两位小数，单位元
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetPerFace(_perFace string) error {
	r._perFace = _perFace
	r.Set("per_face", _perFace)
	return nil
}

// GetPerFace PerFace Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetPerFace() string {
	return r._perFace
}

// SetName is Name Setter
// 淘礼金名称，最大10个字符
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetName() string {
	return r._name
}

// SetItemId is ItemId Setter
// 宝贝ID或营销ID
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetItemId() string {
	return r._itemId
}

// SetCampaignType is CampaignType Setter
// CPS佣金类型
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetCampaignType(_campaignType string) error {
	r._campaignType = _campaignType
	r.Set("campaign_type", _campaignType)
	return nil
}

// GetCampaignType CampaignType Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetCampaignType() string {
	return r._campaignType
}

// SetSecurityLevel is SecurityLevel Setter
// 必须传入0
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetSecurityLevel(_securityLevel int64) error {
	r._securityLevel = _securityLevel
	r.Set("security_level", _securityLevel)
	return nil
}

// GetSecurityLevel SecurityLevel Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetSecurityLevel() int64 {
	return r._securityLevel
}

// SetUseEndTimeMode is UseEndTimeMode Setter
// 结束日期的模式,1:相对时间，2:绝对时间
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetUseEndTimeMode(_useEndTimeMode int64) error {
	r._useEndTimeMode = _useEndTimeMode
	r.Set("use_end_time_mode", _useEndTimeMode)
	return nil
}

// GetUseEndTimeMode UseEndTimeMode Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetUseEndTimeMode() int64 {
	return r._useEndTimeMode
}

// SetUserTotalWinNumLimit is UserTotalWinNumLimit Setter
// 单用户累计中奖次数上限
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetUserTotalWinNumLimit(_userTotalWinNumLimit int64) error {
	r._userTotalWinNumLimit = _userTotalWinNumLimit
	r.Set("user_total_win_num_limit", _userTotalWinNumLimit)
	return nil
}

// GetUserTotalWinNumLimit UserTotalWinNumLimit Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetUserTotalWinNumLimit() int64 {
	return r._userTotalWinNumLimit
}

// SetTotalNum is TotalNum Setter
// 淘礼金总个数
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetTotalNum(_totalNum int64) error {
	r._totalNum = _totalNum
	r.Set("total_num", _totalNum)
	return nil
}

// GetTotalNum TotalNum Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetTotalNum() int64 {
	return r._totalNum
}

// SetAdzoneId is AdzoneId Setter
// 妈妈广告位Id
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSecuritySwitch is SecuritySwitch Setter
// 必须设置为true，默认开启安全
func (r *TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) SetSecuritySwitch(_securitySwitch bool) error {
	r._securitySwitch = _securitySwitch
	r.Set("security_switch", _securitySwitch)
	return nil
}

// GetSecuritySwitch SecuritySwitch Getter
func (r TaobaoTbkDgVegasTljTemporaryCreateAPIRequest) GetSecuritySwitch() bool {
	return r._securitySwitch
}