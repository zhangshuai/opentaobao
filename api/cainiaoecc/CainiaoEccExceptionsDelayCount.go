package cainiaoecc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoecc"
)

// CainiaoEccExceptionsDelayCount 菜鸟控制塔包裹滞留异常统计信息获取
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
func CainiaoEccExceptionsDelayCount(clt *core.SDKClient, req *cainiaoecc.CainiaoEccExceptionsDelayCountAPIRequest, resp *cainiaoecc.CainiaoEccExceptionsDelayCountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
