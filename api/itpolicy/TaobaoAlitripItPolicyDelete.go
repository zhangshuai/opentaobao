package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItPolicyDelete 【国际机票销售规则】单条删除
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
func TaobaoAlitripItPolicyDelete(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyDeleteAPIRequest, resp *itpolicy.TaobaoAlitripItPolicyDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
