package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeAnnouncementPublish 支付宝疫苗POV公告通知
// alibaba.health.vaccin.notice.announcement.publish
//
// 支付宝疫苗POV发布公告提醒信息
func AlibabaHealthVaccinNoticeAnnouncementPublish(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeAnnouncementPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
