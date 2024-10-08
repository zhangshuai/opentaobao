package normalvisa

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// AlitripTravelVisaSignSend 签证批量申请人送签接口
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
func AlitripTravelVisaSignSend(ctx context.Context, clt *core.SDKClient, req *normalvisa.AlitripTravelVisaSignSendAPIRequest, resp *normalvisa.AlitripTravelVisaSignSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
