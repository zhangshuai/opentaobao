package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// WdkRexoutResourceListCheck ReX应用中心资源更新检测-外部
// wdk.rexout.resource.list.check
//
// ReX应用中心资源更新检测-外部
func WdkRexoutResourceListCheck(clt *core.SDKClient, req *util.WdkRexoutResourceListCheckAPIRequest, resp *util.WdkRexoutResourceListCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
