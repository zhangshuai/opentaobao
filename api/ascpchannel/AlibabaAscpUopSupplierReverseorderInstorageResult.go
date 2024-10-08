package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierReverseorderInstorageResult 逆向销退入库单到仓结果回告
// alibaba.ascp.uop.supplier.reverseorder.instorage.result
//
// ERP回告销退入库单到仓信息回告
func AlibabaAscpUopSupplierReverseorderInstorageResult(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
