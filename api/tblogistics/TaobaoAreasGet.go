package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaoareasget 查询地址区域
// taobao.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// &lt;a href=&#34;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&#34; target=&#34;_blank&#34;&gt; http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&lt;/a&gt;
func Taobaoareasget(clt *core.SDKClient, req *tblogistics.TaobaoareasgetAPIRequest, session string) (*tblogistics.TaobaoareasgetAPIResponse, error) {
	var resp tblogistics.TaobaoareasgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}