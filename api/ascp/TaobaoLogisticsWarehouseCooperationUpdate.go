package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehousecooperationupdate 合作商家信息同步
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
func Taobaologisticswarehousecooperationupdate(clt *core.SDKClient, req *ascp.TaobaologisticswarehousecooperationupdateAPIRequest, session string) (*ascp.TaobaologisticswarehousecooperationupdateAPIResponse, error) {
	var resp ascp.TaobaologisticswarehousecooperationupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}