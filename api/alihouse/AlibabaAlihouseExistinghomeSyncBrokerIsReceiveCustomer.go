package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomer 经纪人接待状态变更
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
func AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomer(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
