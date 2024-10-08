package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleReportMediaUpload 验货报告上传文件
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
func AlibabaIdleReportMediaUpload(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleReportMediaUploadAPIRequest, resp *idle.AlibabaIdleReportMediaUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
