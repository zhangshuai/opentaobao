package idleparttime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleparttime"
)

// AlibabaIdleParttimeJobsync 兼职岗位同步
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
func AlibabaIdleParttimeJobsync(clt *core.SDKClient, req *idleparttime.AlibabaIdleParttimeJobsyncAPIRequest, resp *idleparttime.AlibabaIdleParttimeJobsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
