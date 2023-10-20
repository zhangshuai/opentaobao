package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabaiworkcorehrsgetperson 获取神鲸用户基本信息
// alibaba.iwork.core.hrs.getperson
//
// 神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
func Alibabaiworkcorehrsgetperson(clt *core.SDKClient, req *campus.AlibabaiworkcorehrsgetpersonAPIRequest, session string) (*campus.AlibabaiworkcorehrsgetpersonAPIResponse, error) {
	var resp campus.AlibabaiworkcorehrsgetpersonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
