package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalModifyrestime 修改预约时间
// alibaba.alihealth.reserve.dental.modifyrestime
//
// 修改预约时间
func AlibabaAlihealthReserveDentalModifyrestime(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalModifyrestimeAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalModifyrestimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
