package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenVersionApply 数据同步版本号申请
// alibaba.wdk.marketing.open.version.apply
//
// 数据同步版本号申请
func AlibabaWdkMarketingOpenVersionApply(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenVersionApplyAPIRequest, resp *wdk.AlibabaWdkMarketingOpenVersionApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
