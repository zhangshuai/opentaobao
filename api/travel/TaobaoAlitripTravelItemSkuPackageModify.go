package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemSkuPackageModify 【API3.0】套餐级别日历价格库存增删操作
// taobao.alitrip.travel.item.sku.package.modify
//
// 【API3.0】套餐级别日历价格库存增删操作
func TaobaoAlitripTravelItemSkuPackageModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuPackageModifyAPIRequest, resp *travel.TaobaoAlitripTravelItemSkuPackageModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
