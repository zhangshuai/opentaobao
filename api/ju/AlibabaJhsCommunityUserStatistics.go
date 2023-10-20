package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityUserStatistics 聚划算社群用户行为上报
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
func AlibabaJhsCommunityUserStatistics(clt *core.SDKClient, req *ju.AlibabaJhsCommunityUserStatisticsAPIRequest, session string) (*ju.AlibabaJhsCommunityUserStatisticsAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityUserStatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
