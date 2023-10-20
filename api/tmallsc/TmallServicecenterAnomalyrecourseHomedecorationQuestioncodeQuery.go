package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationquestioncodequery 天猫服务平台商家投诉单问题列表查询
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
func Tmallservicecenteranomalyrecoursehomedecorationquestioncodequery(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}