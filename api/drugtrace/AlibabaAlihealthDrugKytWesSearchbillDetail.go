package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchbillDetail 查询单据详情
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
func AlibabaAlihealthDrugKytWesSearchbillDetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
