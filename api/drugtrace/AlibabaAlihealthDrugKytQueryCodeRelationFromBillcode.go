package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQueryCodeRelationFromBillcode 根据单据号码查询码单据详情和码信息
// alibaba.alihealth.drug.kyt.query.code.relation.from.billcode
//
// 根据单据号码查询码单据详情和码信息
func AlibabaAlihealthDrugKytQueryCodeRelationFromBillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
