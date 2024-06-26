package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportDelete 质检报告删除接口
// taobao.qt.report.delete
//
// 删除质检报告
func TaobaoQtReportDelete(clt *core.SDKClient, req *qt.TaobaoQtReportDeleteAPIRequest, resp *qt.TaobaoQtReportDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
