package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreprocessConfirm 仓内加工单确认接口
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
func TaobaoQimenStoreprocessConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenStoreprocessConfirmAPIRequest, resp *qimen.TaobaoQimenStoreprocessConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
