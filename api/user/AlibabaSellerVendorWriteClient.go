package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaSellerVendorWriteClient 客户动态回写
// alibaba.seller.vendor.write.client
//
// 客户动态开放API接口，外部服务商回写数据
func AlibabaSellerVendorWriteClient(ctx context.Context, clt *core.SDKClient, req *user.AlibabaSellerVendorWriteClientAPIRequest, resp *user.AlibabaSellerVendorWriteClientAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
