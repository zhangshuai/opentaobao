package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoWirelessContentCheck 无线开放内容检查
// taobao.wireless.content.check
//
// 无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 &lt;a href=&#34;https://help.aliyun.com/document_detail/70439.html&#34; target=&#34;blank&#34;&gt;阿里云内容安全&lt;/a&gt;
func TaobaoWirelessContentCheck(ctx context.Context, clt *core.SDKClient, req *util.TaobaoWirelessContentCheckAPIRequest, resp *util.TaobaoWirelessContentCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
