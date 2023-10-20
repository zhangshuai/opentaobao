package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabaservicecenterworkcardrelatedskuquery 查询工单关联的服务项
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
func Alibabaservicecenterworkcardrelatedskuquery(clt *core.SDKClient, req *tmallsc.AlibabaservicecenterworkcardrelatedskuqueryAPIRequest, session string) (*tmallsc.AlibabaservicecenterworkcardrelatedskuqueryAPIResponse, error) {
	var resp tmallsc.AlibabaservicecenterworkcardrelatedskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
