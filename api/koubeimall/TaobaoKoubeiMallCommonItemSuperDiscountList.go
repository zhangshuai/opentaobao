package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonItemSuperDiscountList 查询商圈内的超值特惠商品信息
// taobao.koubei.mall.common.item.super.discount.list
//
// 查询商圈超值特惠商品信息列表
func TaobaoKoubeiMallCommonItemSuperDiscountList(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
