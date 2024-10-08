package ioti

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslEslimageSendimage 下发厂测初始化图片
// alibaba.it.esl.eslimage.sendimage
//
// 工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
func AlibabaItEslEslimageSendimage(ctx context.Context, clt *core.SDKClient, req *ioti.AlibabaItEslEslimageSendimageAPIRequest, resp *ioti.AlibabaItEslEslimageSendimageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
