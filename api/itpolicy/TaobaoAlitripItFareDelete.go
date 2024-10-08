package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItFareDelete 【国际机票自有政策】单条删除
// taobao.alitrip.it.fare.delete
//
// 自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
func TaobaoAlitripItFareDelete(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareDeleteAPIRequest, resp *itpolicy.TaobaoAlitripItFareDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
