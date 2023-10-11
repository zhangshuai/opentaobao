package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenSingleitemSynchronize 商品同步接口
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
func TaobaoQimenSingleitemSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenSingleitemSynchronizeAPIRequest, session string) (*qimen.TaobaoQimenSingleitemSynchronizeAPIResponse, error) {
	var resp qimen.TaobaoQimenSingleitemSynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
