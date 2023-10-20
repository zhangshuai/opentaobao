package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterEntityDelete 删除外部成本中心人员信息
// alitrip.btrip.cost.center.entity.delete
//
// 删除外部成本中心人员信息
func AlitripBtripCostCenterEntityDelete(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntityDeleteAPIRequest, resp *btrip.AlitripBtripCostCenterEntityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
