package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivitySkuAdd 特价活动新增商品
// alibaba.retail.marketing.itemdiscount.activity.sku.add
//
// 新增或更新活动商品信息【同城零售】
func AlibabaRetailMarketingItemdiscountActivitySkuAdd(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
