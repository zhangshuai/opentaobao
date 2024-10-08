package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuNumberPut ISV上传数据接口
// taobao.qianniu.number.put
//
// ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。
func TaobaoQianniuNumberPut(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuNumberPutAPIRequest, resp *qianniu.TaobaoQianniuNumberPutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
