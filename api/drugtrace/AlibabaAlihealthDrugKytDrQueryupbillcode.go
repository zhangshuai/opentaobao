package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrqueryupbillcode 查询上游企业出库单据号
// alibaba.alihealth.drug.kyt.dr.queryupbillcode
//
// 疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号
func Alibabaalihealthdrugkytdrqueryupbillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrqueryupbillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrqueryupbillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrqueryupbillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
