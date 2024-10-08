package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceProductSearch 搜索保险产品
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
func AlitripFlightInsuranceProductSearch(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceProductSearchAPIRequest, resp *flightuppc.AlitripFlightInsuranceProductSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
