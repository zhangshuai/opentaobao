package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmosfundcancelbill 取消付款单
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
func Alibabamjmosfundcancelbill(clt *core.SDKClient, req *mos.AlibabamjmosfundcancelbillAPIRequest, session string) (*mos.AlibabamjmosfundcancelbillAPIResponse, error) {
	var resp mos.AlibabamjmosfundcancelbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
