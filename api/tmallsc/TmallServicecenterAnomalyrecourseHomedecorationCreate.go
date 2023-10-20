package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationcreate 天猫服务平台服务商代商家发起投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
func Tmallservicecenteranomalyrecoursehomedecorationcreate(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationcreateAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationcreateAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}