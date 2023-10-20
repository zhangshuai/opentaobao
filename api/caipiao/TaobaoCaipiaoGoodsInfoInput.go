package caipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// Taobaocaipiaogoodsinfoinput 录入参加送彩票商品信息
// taobao.caipiao.goods.info.input
//
// 录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
func Taobaocaipiaogoodsinfoinput(clt *core.SDKClient, req *caipiao.TaobaocaipiaogoodsinfoinputAPIRequest, session string) (*caipiao.TaobaocaipiaogoodsinfoinputAPIResponse, error) {
	var resp caipiao.TaobaocaipiaogoodsinfoinputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
