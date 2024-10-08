package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtOpentradeGetproductinfo 查询新虚拟产品配置信息
// alibaba.alicom.vt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
func AlibabaAlicomVtOpentradeGetproductinfo(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomVtOpentradeGetproductinfoAPIRequest, resp *alicom.AlibabaAlicomVtOpentradeGetproductinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
