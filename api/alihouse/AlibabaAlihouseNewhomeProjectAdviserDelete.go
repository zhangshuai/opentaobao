package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectAdviserDelete 删除楼盘顾问
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
func AlibabaAlihouseNewhomeProjectAdviserDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
