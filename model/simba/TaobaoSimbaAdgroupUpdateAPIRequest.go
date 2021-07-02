package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupUpdateAPIRequest 更新一个推广组的信息 API请求
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
type TaobaoSimbaAdgroupUpdateAPIRequest struct {
	model.Params
	// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
	_useNonsearchDefaultPrice string
	// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
	_onlineStatus string
	// 推广组Id
	_adgroupId int64
	// 默认出价，单位是分，不能小于5
	_defaultPrice int64
	// 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
	_nonsearchMaxPrice int64
	// 主人昵称
	_nick string
}

// NewTaobaoSimbaAdgroupUpdateRequest 初始化TaobaoSimbaAdgroupUpdateAPIRequest对象
func NewTaobaoSimbaAdgroupUpdateRequest() *TaobaoSimbaAdgroupUpdateAPIRequest {
	return &TaobaoSimbaAdgroupUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UseNonsearchDefaultPrice Setter
// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetUseNonsearchDefaultPrice(_useNonsearchDefaultPrice string) error {
	r._useNonsearchDefaultPrice = _useNonsearchDefaultPrice
	r.Set("use_nonsearch_default_price", _useNonsearchDefaultPrice)
	return nil
}

// Get UseNonsearchDefaultPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetUseNonsearchDefaultPrice() string {
	return r._useNonsearchDefaultPrice
}

// Set is OnlineStatus Setter
// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetOnlineStatus(_onlineStatus string) error {
	r._onlineStatus = _onlineStatus
	r.Set("online_status", _onlineStatus)
	return nil
}

// Get OnlineStatus Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetOnlineStatus() string {
	return r._onlineStatus
}

// Set is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is DefaultPrice Setter
// 默认出价，单位是分，不能小于5
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetDefaultPrice(_defaultPrice int64) error {
	r._defaultPrice = _defaultPrice
	r.Set("default_price", _defaultPrice)
	return nil
}

// Get DefaultPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetDefaultPrice() int64 {
	return r._defaultPrice
}

// Set is NonsearchMaxPrice Setter
// 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetNonsearchMaxPrice(_nonsearchMaxPrice int64) error {
	r._nonsearchMaxPrice = _nonsearchMaxPrice
	r.Set("nonsearch_max_price", _nonsearchMaxPrice)
	return nil
}

// Get NonsearchMaxPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetNonsearchMaxPrice() int64 {
	return r._nonsearchMaxPrice
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetNick() string {
	return r._nick
}
