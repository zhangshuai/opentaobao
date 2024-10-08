package aliyun

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// PushAliyuncsComPush20150518 云推送指令API
// push.aliyuncs.com.push.20150518
//
// 阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
func PushAliyuncsComPush20150518(ctx context.Context, clt *core.SDKClient, req *aliyun.PushAliyuncsComPush20150518APIRequest, resp *aliyun.PushAliyuncsComPush20150518APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
