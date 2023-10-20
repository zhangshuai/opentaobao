package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpstart 聚安全实人认证开始
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
func Alibabasecurityjaqrpstart(clt *core.SDKClient, req *security.AlibabasecurityjaqrpstartAPIRequest, session string) (*security.AlibabasecurityjaqrpstartAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpstartAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
