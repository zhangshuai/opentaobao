package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// Taobaobrandstartshoprptadgroupget 明星店铺推广单元报表数据查询
// taobao.brand.startshop.rpt.adgroup.get
//
// 获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func Taobaobrandstartshoprptadgroupget(clt *core.SDKClient, req *brandhub.TaobaobrandstartshoprptadgroupgetAPIRequest, session string) (*brandhub.TaobaobrandstartshoprptadgroupgetAPIResponse, error) {
	var resp brandhub.TaobaobrandstartshoprptadgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
