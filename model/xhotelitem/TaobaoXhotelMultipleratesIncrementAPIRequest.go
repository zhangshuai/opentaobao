package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultipleratesIncrementAPIRequest 复杂房价推送接口（批量增量） API请求
// taobao.xhotel.multiplerates.increment
//
// 复杂房价批量增量更新，只会更新指定日期的信息
// 完全涵盖了taobao.xhotel.rates.increment接口的功能
type TaobaoXhotelMultipleratesIncrementAPIRequest struct {
	model.Params
	// 批量全量修改价格和库存信息,会以请求参数中的数据覆盖掉原来报价库存数据。A:useRoomInventory:是否使用room级别共享库存，可选值 true false  2、false时：使用rate级别私有库存，此时如果填写了库存，那么会写入库存表。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:basePrice 基本价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存） 支持状态库存， 60000(状态库存关) 61000(状态库存开);E:occupancy为入住人数，范围为1~10;F:lengthofStay为连住天数，范围为1~10；G:taxAndFee为总税费；H:addBedPrice为加床价；I:addPersonPrice为加人价；J:rateSwitch为开关房状态，1为开房，0为关房。K:支持outRoomId和ratePlanCode来更新报价库存。L:childnum为儿童人数。M:infantnum为婴儿人数。N:ckinSwitch为入住开关(0，关闭；1，打开) O：ckoutSwitch为离店开关 (0，关闭；1，打开) P:lockStartTime锁库存开始时间 Q:lockEndTime锁库存截止时间
	_rateQuotaMap string
}

// NewTaobaoXhotelMultipleratesIncrementRequest 初始化TaobaoXhotelMultipleratesIncrementAPIRequest对象
func NewTaobaoXhotelMultipleratesIncrementRequest() *TaobaoXhotelMultipleratesIncrementAPIRequest {
	return &TaobaoXhotelMultipleratesIncrementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.multiplerates.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RateQuotaMap Setter
// 批量全量修改价格和库存信息,会以请求参数中的数据覆盖掉原来报价库存数据。A:useRoomInventory:是否使用room级别共享库存，可选值 true false  2、false时：使用rate级别私有库存，此时如果填写了库存，那么会写入库存表。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:basePrice 基本价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存） 支持状态库存， 60000(状态库存关) 61000(状态库存开);E:occupancy为入住人数，范围为1~10;F:lengthofStay为连住天数，范围为1~10；G:taxAndFee为总税费；H:addBedPrice为加床价；I:addPersonPrice为加人价；J:rateSwitch为开关房状态，1为开房，0为关房。K:支持outRoomId和ratePlanCode来更新报价库存。L:childnum为儿童人数。M:infantnum为婴儿人数。N:ckinSwitch为入住开关(0，关闭；1，打开) O：ckoutSwitch为离店开关 (0，关闭；1，打开) P:lockStartTime锁库存开始时间 Q:lockEndTime锁库存截止时间
func (r *TaobaoXhotelMultipleratesIncrementAPIRequest) SetRateQuotaMap(_rateQuotaMap string) error {
	r._rateQuotaMap = _rateQuotaMap
	r.Set("rate_quota_map", _rateQuotaMap)
	return nil
}

// Get RateQuotaMap Getter
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetRateQuotaMap() string {
	return r._rateQuotaMap
}
