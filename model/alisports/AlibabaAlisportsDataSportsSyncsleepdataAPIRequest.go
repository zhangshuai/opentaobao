package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncsleepdataAPIRequest 阿里体育数据中心用户睡眠数据同步接口 API请求
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
type AlibabaAlisportsDataSportsSyncsleepdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 入睡时间，格式：y-m-d h:i:s
	_stime string
	// 清醒时间，单位：小时
	_soberTime string
	// 浅度睡眠时间，单位：小时
	_shallowTime string
	// 深度睡眠时间，单位：小时
	_deepTime string
	// 睡眠总时间，单位：小时
	_allTime string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 醒来时间，格式：y-m-d h:i:s
	_etime string
}

// NewAlibabaAlisportsDataSportsSyncsleepdataRequest 初始化AlibabaAlisportsDataSportsSyncsleepdataAPIRequest对象
func NewAlibabaAlisportsDataSportsSyncsleepdataRequest() *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest {
	return &AlibabaAlisportsDataSportsSyncsleepdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncsleepdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// Get AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// Set is Stime Setter
// 入睡时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetStime(_stime string) error {
	r._stime = _stime
	r.Set("stime", _stime)
	return nil
}

// Get Stime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetStime() string {
	return r._stime
}

// Set is SoberTime Setter
// 清醒时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetSoberTime(_soberTime string) error {
	r._soberTime = _soberTime
	r.Set("sober_time", _soberTime)
	return nil
}

// Get SoberTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetSoberTime() string {
	return r._soberTime
}

// Set is ShallowTime Setter
// 浅度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetShallowTime(_shallowTime string) error {
	r._shallowTime = _shallowTime
	r.Set("shallow_time", _shallowTime)
	return nil
}

// Get ShallowTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetShallowTime() string {
	return r._shallowTime
}

// Set is DeepTime Setter
// 深度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetDeepTime(_deepTime string) error {
	r._deepTime = _deepTime
	r.Set("deep_time", _deepTime)
	return nil
}

// Get DeepTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetDeepTime() string {
	return r._deepTime
}

// Set is AllTime Setter
// 睡眠总时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetAllTime(_allTime string) error {
	r._allTime = _allTime
	r.Set("all_time", _allTime)
	return nil
}

// Get AllTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetAllTime() string {
	return r._allTime
}

// Set is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// Get Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// Set is AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// Get AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// Set is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// Get AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// Set is Etime Setter
// 醒来时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// Get Etime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataAPIRequest) GetEtime() string {
	return r._etime
}
