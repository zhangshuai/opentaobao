package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolDeleteactivity 删除商品池活动
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
func AlibabaHmMarketingItempoolDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolDeleteactivityAPIRequest, resp *wdk.AlibabaHmMarketingItempoolDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
