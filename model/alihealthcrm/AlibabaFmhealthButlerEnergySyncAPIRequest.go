package alihealthcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthButlerEnergySyncAPIRequest 同步用户消耗能量 API请求
// alibaba.fmhealth.butler.energy.sync
//
// 同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
type AlibabaFmhealthButlerEnergySyncAPIRequest struct {
	model.Params
	// “S”- s点 “CAL”- 卡路里
	_energyType string
	// 阿里用户id
	_userId int64
	// 每日已消耗能量
	_value float64
	// 每日可消耗能量
	_target float64
	// 每日运动消耗能量值
	_sport float64
}

// NewAlibabaFmhealthButlerEnergySyncRequest 初始化AlibabaFmhealthButlerEnergySyncAPIRequest对象
func NewAlibabaFmhealthButlerEnergySyncRequest() *AlibabaFmhealthButlerEnergySyncAPIRequest {
	return &AlibabaFmhealthButlerEnergySyncAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) Reset() {
	r._energyType = ""
	r._userId = 0
	r._value = 0
	r._target = 0
	r._sport = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.fmhealth.butler.energy.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEnergyType is EnergyType Setter
// “S”- s点 “CAL”- 卡路里
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetEnergyType(_energyType string) error {
	r._energyType = _energyType
	r.Set("energy_type", _energyType)
	return nil
}

// GetEnergyType EnergyType Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetEnergyType() string {
	return r._energyType
}

// SetUserId is UserId Setter
// 阿里用户id
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetValue is Value Setter
// 每日已消耗能量
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetValue(_value float64) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetValue() float64 {
	return r._value
}

// SetTarget is Target Setter
// 每日可消耗能量
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetTarget(_target float64) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetTarget() float64 {
	return r._target
}

// SetSport is Sport Setter
// 每日运动消耗能量值
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetSport(_sport float64) error {
	r._sport = _sport
	r.Set("sport", _sport)
	return nil
}

// GetSport Sport Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetSport() float64 {
	return r._sport
}

var poolAlibabaFmhealthButlerEnergySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFmhealthButlerEnergySyncRequest()
	},
}

// GetAlibabaFmhealthButlerEnergySyncRequest 从 sync.Pool 获取 AlibabaFmhealthButlerEnergySyncAPIRequest
func GetAlibabaFmhealthButlerEnergySyncAPIRequest() *AlibabaFmhealthButlerEnergySyncAPIRequest {
	return poolAlibabaFmhealthButlerEnergySyncAPIRequest.Get().(*AlibabaFmhealthButlerEnergySyncAPIRequest)
}

// ReleaseAlibabaFmhealthButlerEnergySyncAPIRequest 将 AlibabaFmhealthButlerEnergySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaFmhealthButlerEnergySyncAPIRequest(v *AlibabaFmhealthButlerEnergySyncAPIRequest) {
	v.Reset()
	poolAlibabaFmhealthButlerEnergySyncAPIRequest.Put(v)
}
