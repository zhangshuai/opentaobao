package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupappointmentorderdifferencedetail 预约单差异明细查询
// cainiao.global.im.pickup.appointment.order.difference.detail
//
// 预约单差异明细查询
func Cainiaoglobalimpickupappointmentorderdifferencedetail(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupappointmentorderdifferencedetailAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupappointmentorderdifferencedetailAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupappointmentorderdifferencedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}