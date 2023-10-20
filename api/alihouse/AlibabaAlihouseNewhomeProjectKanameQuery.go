package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectKanameQuery 查询KA楼盘名称
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
func AlibabaAlihouseNewhomeProjectKanameQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
