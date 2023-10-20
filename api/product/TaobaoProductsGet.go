package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoproductsget 获取产品列表
// taobao.products.get
//
// 根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
// 注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页
func Taobaoproductsget(clt *core.SDKClient, req *product.TaobaoproductsgetAPIRequest, session string) (*product.TaobaoproductsgetAPIResponse, error) {
	var resp product.TaobaoproductsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
