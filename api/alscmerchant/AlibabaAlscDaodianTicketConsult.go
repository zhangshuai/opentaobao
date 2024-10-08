package alscmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscDaodianTicketConsult 券码预览
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
func AlibabaAlscDaodianTicketConsult(ctx context.Context, clt *core.SDKClient, req *alscmerchant.AlibabaAlscDaodianTicketConsultAPIRequest, resp *alscmerchant.AlibabaAlscDaodianTicketConsultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
