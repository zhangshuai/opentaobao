package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallback 每日优鲜仓作业单回传接口
// alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback
//
// 家乐福仓作业单回传接口
func AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
