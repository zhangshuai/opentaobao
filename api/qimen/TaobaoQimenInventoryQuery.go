package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryQuery 库存查询接口（多商品）
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
func TaobaoQimenInventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryQueryAPIRequest, resp *qimen.TaobaoQimenInventoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
