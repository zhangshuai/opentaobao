package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomepicturesync 图片数据同步
// alibaba.alihouse.newhome.picture.sync
//
// 图片数据同步
func Alibabaalihousenewhomepicturesync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomepicturesyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomepicturesyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomepicturesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
