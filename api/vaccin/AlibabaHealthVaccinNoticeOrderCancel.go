package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeOrderCancel 福州疫苗取消预约
// alibaba.health.vaccin.notice.order.cancel
//
// 福州疫苗用户取消预约接口
func AlibabaHealthVaccinNoticeOrderCancel(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderCancelAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
