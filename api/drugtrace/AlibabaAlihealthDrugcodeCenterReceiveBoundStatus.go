package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeCenterReceiveBoundStatus 接收中央随机化系统和临床研究机构的绑定确认状态
// alibaba.alihealth.drugcode.center.receive.bound.status
//
// 临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
func AlibabaAlihealthDrugcodeCenterReceiveBoundStatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
