package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentHandrefundRefundfee 代理商手动退款接口
// taobao.train.agent.handrefund.refundfee
//
// 火车票代理商手动退款接口
func TaobaoTrainAgentHandrefundRefundfee(clt *core.SDKClient, req *train.TaobaoTrainAgentHandrefundRefundfeeAPIRequest, resp *train.TaobaoTrainAgentHandrefundRefundfeeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
