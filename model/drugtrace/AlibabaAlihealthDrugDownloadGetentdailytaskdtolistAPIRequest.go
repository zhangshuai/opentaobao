package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest 码上放心数据落地-获取每天日报 API请求
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
type AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest struct {
	model.Params
	// appkey
	_appKeyN string
	// 统计的开始时间
	_startTime string
	// 统计的结束时间
	_endTime string
}

// NewAlibabaalihealthdrugdownloadgetentdailytaskdtolistRequest 初始化AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest对象
func NewAlibabaalihealthdrugdownloadgetentdailytaskdtolistRequest() *AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest {
	return &AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.getentdailytaskdtolist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appkey
func (r *AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetStartTime is StartTime Setter
// 统计的开始时间
func (r *AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 统计的结束时间
func (r *AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest) GetEndTime() string {
	return r._endTime
}
