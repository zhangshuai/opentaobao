package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkCouponTemplateQueryumpactid 通过券模板查询券活动id接口
// alibaba.wdk.coupon.template.queryumpactid
//
// 当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
func AlibabaWdkCouponTemplateQueryumpactid(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkCouponTemplateQueryumpactidAPIRequest, resp *wdk.AlibabaWdkCouponTemplateQueryumpactidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
