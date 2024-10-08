package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiProduct 商家查询物流商产品类型接口
// cainiao.waybill.ii.product
//
// 商家可以查询物流商的产品类型和服务能力。
func CainiaoWaybillIiProduct(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiProductAPIRequest, resp *waybill.CainiaoWaybillIiProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
