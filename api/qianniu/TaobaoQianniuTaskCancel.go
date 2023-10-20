package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskCancel 取消轻任务
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
func TaobaoQianniuTaskCancel(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskCancelAPIRequest, resp *qianniu.TaobaoQianniuTaskCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
