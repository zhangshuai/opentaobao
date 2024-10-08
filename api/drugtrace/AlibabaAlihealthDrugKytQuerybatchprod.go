package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerybatchprod 批次产品查询(根据企业名和批次号查询产品信息)
// alibaba.alihealth.drug.kyt.querybatchprod
//
// 根据企业名和批次号查询药品信息，支持使用更名之前的老企业名查询，支持批次号大小写模糊，应用于药店或医院入库环节，通过在入库环节获取赋码的产品目录，可强制要求对相应的产品必须进行扫码入库；
func AlibabaAlihealthDrugKytQuerybatchprod(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerybatchprodAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerybatchprodAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
