package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiPriceGet 获取价格预估信息
// alibaba.happytrip.taxi.price.get
//
// 打车价格预估
func AlibabaHappytripTaxiPriceGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiPriceGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiPriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
