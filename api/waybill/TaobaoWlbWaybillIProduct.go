package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Taobaowlbwaybilliproduct 商家查询物流商产品类型接口
// taobao.wlb.waybill.i.product
//
// 商家可以查询物流商的产品类型和服务能力。
func Taobaowlbwaybilliproduct(clt *core.SDKClient, req *waybill.TaobaowlbwaybilliproductAPIRequest, session string) (*waybill.TaobaowlbwaybilliproductAPIResponse, error) {
	var resp waybill.TaobaowlbwaybilliproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
