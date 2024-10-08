package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaVoiceNumDoublecall 聚石塔语音双呼接口
// alibaba.aliqin.ta.voice.num.doublecall
//
// 根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
func AlibabaAliqinTaVoiceNumDoublecall(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinTaVoiceNumDoublecallAPIRequest, resp *alicom.AlibabaAliqinTaVoiceNumDoublecallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
