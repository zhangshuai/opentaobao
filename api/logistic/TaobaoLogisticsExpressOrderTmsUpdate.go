package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderTmsUpdate 服务商修改上门取退时间接口
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
func TaobaoLogisticsExpressOrderTmsUpdate(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderTmsUpdateAPIRequest, session string) (*logistic.TaobaoLogisticsExpressOrderTmsUpdateAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressOrderTmsUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
