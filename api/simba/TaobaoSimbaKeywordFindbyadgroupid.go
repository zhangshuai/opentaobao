package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordFindbyadgroupid 根据推广单元id获取关键词
// taobao.simba.keyword.findbyadgroupid
//
// 根据一个关键词Id列表取得一组关键词
func TaobaoSimbaKeywordFindbyadgroupid(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordFindbyadgroupidAPIRequest, resp *simba.TaobaoSimbaKeywordFindbyadgroupidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
