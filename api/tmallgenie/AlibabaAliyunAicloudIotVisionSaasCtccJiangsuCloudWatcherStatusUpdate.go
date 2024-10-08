package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdate 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
func AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdate(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest, resp *tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
