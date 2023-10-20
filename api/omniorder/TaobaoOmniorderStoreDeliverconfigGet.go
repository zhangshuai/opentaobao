package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoredeliverconfigget 查询门店发货配置内容
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
func Taobaoomniorderstoredeliverconfigget(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoredeliverconfiggetAPIRequest, session string) (*omniorder.TaobaoomniorderstoredeliverconfiggetAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoredeliverconfiggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
