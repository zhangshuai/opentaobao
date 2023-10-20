package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductschemamatch product匹配接口
// tmall.product.schema.match
//
// 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
func Tmallproductschemamatch(clt *core.SDKClient, req *tbitem.TmallproductschemamatchAPIRequest, session string) (*tbitem.TmallproductschemamatchAPIResponse, error) {
	var resp tbitem.TmallproductschemamatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}