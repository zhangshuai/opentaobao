package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// Alibabafootscanminireportfragmentfirst 第一只脚生成报告接口
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
func Alibabafootscanminireportfragmentfirst(clt *core.SDKClient, req *foodscan.AlibabafootscanminireportfragmentfirstAPIRequest, session string) (*foodscan.AlibabafootscanminireportfragmentfirstAPIResponse, error) {
	var resp foodscan.AlibabafootscanminireportfragmentfirstAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
