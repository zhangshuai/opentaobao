package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomWttOpentradeGetproductinfo 查询存送产品信息
// alibaba.alicom.wtt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
func AlibabaAlicomWttOpentradeGetproductinfo(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomWttOpentradeGetproductinfoAPIRequest, resp *alicom.AlibabaAlicomWttOpentradeGetproductinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
