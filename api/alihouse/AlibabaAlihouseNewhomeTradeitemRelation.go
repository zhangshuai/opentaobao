package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemRelation 货独立绑定货品
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
func AlibabaAlihouseNewhomeTradeitemRelation(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
