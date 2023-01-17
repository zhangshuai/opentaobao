package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameScoreReportAPIRequest 云游戏战绩上传通用接口 API请求
// alibaba.cgame.score.report
//
// 阿里云游戏, CP游戏合作方通用游戏结果回传接口
type AlibabaCgameScoreReportAPIRequest struct {
	model.Params
	// 通用战绩回传数据
	_reportData *CpCallbackReportDto
}

// NewAlibabaCgameScoreReportRequest 初始化AlibabaCgameScoreReportAPIRequest对象
func NewAlibabaCgameScoreReportRequest() *AlibabaCgameScoreReportAPIRequest {
	return &AlibabaCgameScoreReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameScoreReportAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.score.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameScoreReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameScoreReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportData is ReportData Setter
// 通用战绩回传数据
func (r *AlibabaCgameScoreReportAPIRequest) SetReportData(_reportData *CpCallbackReportDto) error {
	r._reportData = _reportData
	r.Set("report_data", _reportData)
	return nil
}

// GetReportData ReportData Getter
func (r AlibabaCgameScoreReportAPIRequest) GetReportData() *CpCallbackReportDto {
	return r._reportData
}
