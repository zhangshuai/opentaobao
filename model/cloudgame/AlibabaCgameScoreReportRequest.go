package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏战绩上传通用接口 API请求
alibaba.cgame.score.report

阿里云游戏, CP游戏合作方通用游戏结果回传接口
*/
type AlibabaCgameScoreReportRequest struct {
    model.Params
    // 通用战绩回传数据
    reportData   *CpCallbackReportDto
}

// 初始化AlibabaCgameScoreReportRequest对象
func NewAlibabaCgameScoreReportRequest() *AlibabaCgameScoreReportRequest{
    return &AlibabaCgameScoreReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCgameScoreReportRequest) GetApiMethodName() string {
    return "alibaba.cgame.score.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCgameScoreReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportData Setter
// 通用战绩回传数据
func (r *AlibabaCgameScoreReportRequest) SetReportData(reportData *CpCallbackReportDto) error {
    r.reportData = reportData
    r.Set("report_data", reportData)
    return nil
}

// ReportData Getter
func (r AlibabaCgameScoreReportRequest) GetReportData() *CpCallbackReportDto {
    return r.reportData
}