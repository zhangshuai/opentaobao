package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyDrugcodes 查询药品是否赋码
// alibaba.alihealth.drug.kyt.yy.drugcodes
//
// 药品是否赋码
func AlibabaAlihealthDrugKytYyDrugcodes(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyDrugcodesAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyDrugcodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
