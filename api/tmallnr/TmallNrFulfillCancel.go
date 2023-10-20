package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillCancel 取消上门揽件
// tmall.nr.fulfill.cancel
//
// 新零售门店业务取消上门揽件
func TmallNrFulfillCancel(clt *core.SDKClient, req *tmallnr.TmallNrFulfillCancelAPIRequest, resp *tmallnr.TmallNrFulfillCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
