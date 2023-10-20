package beehive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/beehive"
)

// TaobaoBeehiveItemCpsUrl 分佣链接生成接口
// taobao.beehive.item.cps.url
//
// 传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
func TaobaoBeehiveItemCpsUrl(clt *core.SDKClient, req *beehive.TaobaoBeehiveItemCpsUrlAPIRequest, resp *beehive.TaobaoBeehiveItemCpsUrlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
