package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// AlibabaKclubKcQaSearch 知识云-知识检索
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
func AlibabaKclubKcQaSearch(clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchAPIRequest, resp *kclub.AlibabaKclubKcQaSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
