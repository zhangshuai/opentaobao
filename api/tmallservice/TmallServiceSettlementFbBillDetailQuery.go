package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettlementFbBillDetailQuery 服务商工单结算对账查询-流水查询
// tmall.service.settlement.fb.bill.detail.query
//
// 服务商工单结算对账查询-流水查询，用于查询服务工单费用流水，含服务费、退款、分成、提现等。
func TmallServiceSettlementFbBillDetailQuery(clt *core.SDKClient, req *tmallservice.TmallServiceSettlementFbBillDetailQueryAPIRequest, session string) (*tmallservice.TmallServiceSettlementFbBillDetailQueryAPIResponse, error) {
	var resp tmallservice.TmallServiceSettlementFbBillDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
