package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxianginventorybatchquery 批量查询库存
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
func Alibabadchainaoxianginventorybatchquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxianginventorybatchqueryAPIRequest, session string) (*ascp.AlibabadchainaoxianginventorybatchqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxianginventorybatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
