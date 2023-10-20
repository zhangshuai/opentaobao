package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaojstastrolabestoreinventoryquery 后端商品库存查询接口
// taobao.jst.astrolabe.storeinventory.query
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
func Taobaojstastrolabestoreinventoryquery(clt *core.SDKClient, req *omniorder.TaobaojstastrolabestoreinventoryqueryAPIRequest, session string) (*omniorder.TaobaojstastrolabestoreinventoryqueryAPIResponse, error) {
	var resp omniorder.TaobaojstastrolabestoreinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
