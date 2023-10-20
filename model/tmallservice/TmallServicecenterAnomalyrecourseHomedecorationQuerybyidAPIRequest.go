package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest 天猫服务平台服务商查询商家投诉单 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
type TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest struct {
	model.Params
	// 投诉单id
	_anomalyRecourseId int64
}

// NewTmallServicecenterAnomalyrecourseHomedecorationQuerybyidRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationQuerybyidRequest() *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAnomalyRecourseId is AnomalyRecourseId Setter
// 投诉单id
func (r *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest) SetAnomalyRecourseId(_anomalyRecourseId int64) error {
	r._anomalyRecourseId = _anomalyRecourseId
	r.Set("anomaly_recourse_id", _anomalyRecourseId)
	return nil
}

// GetAnomalyRecourseId AnomalyRecourseId Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest) GetAnomalyRecourseId() int64 {
	return r._anomalyRecourseId
}
