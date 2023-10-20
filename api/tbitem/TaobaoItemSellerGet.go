package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemsellerget 获取单个商品详细信息
// taobao.item.seller.get
//
// 获取单个商品的全部信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func Taobaoitemsellerget(clt *core.SDKClient, req *tbitem.TaobaoitemsellergetAPIRequest, session string) (*tbitem.TaobaoitemsellergetAPIResponse, error) {
	var resp tbitem.TaobaoitemsellergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}