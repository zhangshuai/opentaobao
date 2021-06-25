package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询质检报告 APIResponse
taobao.qt.reports.get

批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
*/
type TaobaoQtReportsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQtReportsGetResponse `json:"taobao_qt_reports_get_response,omitempty"`
}

type TaobaoQtReportsGetResponse struct {

    // 质检报告列表
    Reports   []QtReport `json:"reports,omitempty"`

}