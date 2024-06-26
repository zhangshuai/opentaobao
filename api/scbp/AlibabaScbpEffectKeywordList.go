package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectKeywordList 关键词报表
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
func AlibabaScbpEffectKeywordList(clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordListAPIRequest, resp *scbp.AlibabaScbpEffectKeywordListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
