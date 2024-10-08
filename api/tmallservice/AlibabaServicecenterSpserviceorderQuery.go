package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterSpserviceorderQuery 服务供应链服务单查询
// alibaba.servicecenter.spserviceorder.query
//
// 服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
func AlibabaServicecenterSpserviceorderQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterSpserviceorderQueryAPIRequest, resp *tmallservice.AlibabaServicecenterSpserviceorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
