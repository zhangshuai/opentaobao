package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinMatchOn isv自主上下线疫苗，可以选择上线还是下线
// alibaba.health.vaccin.match.on
//
// isv自主上下线疫苗，可以选择上线还是下线
func AlibabaHealthVaccinMatchOn(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinMatchOnAPIRequest, session string) (*vaccin.AlibabaHealthVaccinMatchOnAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinMatchOnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
