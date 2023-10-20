package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemjointimg 商品关联子图
// taobao.item.joint.img
//
// * 关联一张商品图片到num_iid指定的商品中&lt;br/&gt;    * 传入的num_iid所对应的商品必须属于当前会话的用户&lt;br/&gt;    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行&lt;br/&gt;    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
func Taobaoitemjointimg(clt *core.SDKClient, req *tbitem.TaobaoitemjointimgAPIRequest, session string) (*tbitem.TaobaoitemjointimgAPIResponse, error) {
	var resp tbitem.TaobaoitemjointimgAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}