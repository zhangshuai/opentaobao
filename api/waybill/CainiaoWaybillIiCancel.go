package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiCancel 商家取消获取的电子面单号
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
func CainiaoWaybillIiCancel(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiCancelAPIRequest, resp *waybill.CainiaoWaybillIiCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
