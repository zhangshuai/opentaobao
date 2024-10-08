package filmtfavatar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSalePaymentQueryVii 获取影院卖品账单--支付账单-V2版本(正逆分离)
// taobao.film.tfavatar.bill.sale.payment.query.vii
//
// 获取影院卖品账单--支付账单-V2版本(正逆分离)
func TaobaoFilmTfavatarBillSalePaymentQueryVii(ctx context.Context, clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
