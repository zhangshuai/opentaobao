package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoBillBookBillsGet 查询虚拟账户明细数据(自研发商家专用)
// taobao.bill.book.bills.get
//
// 查询虚拟账户明细数据
func TaobaoBillBookBillsGet(clt *core.SDKClient, req *bill.TaobaoBillBookBillsGetAPIRequest, resp *bill.TaobaoBillBookBillsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
