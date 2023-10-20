package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemhscodeauditresultsquery 商品hscode信息审核状态查询接口
// tmall.item.hscode.audit.results.query
//
// 通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
func Tmallitemhscodeauditresultsquery(clt *core.SDKClient, req *tbitem.TmallitemhscodeauditresultsqueryAPIRequest, session string) (*tbitem.TmallitemhscodeauditresultsqueryAPIResponse, error) {
	var resp tbitem.TmallitemhscodeauditresultsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
