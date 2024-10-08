package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiDriverLocationGet 司机位置
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
func AlibabaHappytripTaxiDriverLocationGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverLocationGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiDriverLocationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
