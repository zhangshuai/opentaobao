package iotticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iotticket"
)

// CainiaoIotTicketSpComment IoT售后服务商工单备注
// cainiao.iot.ticket.sp.comment
//
// IoT售后服务商工单备注
func CainiaoIotTicketSpComment(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpCommentAPIRequest, resp *iotticket.CainiaoIotTicketSpCommentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
