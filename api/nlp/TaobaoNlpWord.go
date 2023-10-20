package nlp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlp"
)

// Taobaonlpword 文本语言词法分析
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
func Taobaonlpword(clt *core.SDKClient, req *nlp.TaobaonlpwordAPIRequest, session string) (*nlp.TaobaonlpwordAPIResponse, error) {
	var resp nlp.TaobaonlpwordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
