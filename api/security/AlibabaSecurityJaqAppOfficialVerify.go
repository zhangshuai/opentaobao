package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppOfficialVerify 聚安全验证官方应用接口
// alibaba.security.jaq.app.official.verify
//
// 接入用户来查询应用是否为官方应用
func AlibabaSecurityJaqAppOfficialVerify(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppOfficialVerifyAPIRequest, resp *security.AlibabaSecurityJaqAppOfficialVerifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
