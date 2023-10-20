package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemEditSchemaGet 商品编辑获取schema信息
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
func AlibabaItemEditSchemaGet(clt *core.SDKClient, req *tbitem.AlibabaItemEditSchemaGetAPIRequest, resp *tbitem.AlibabaItemEditSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
