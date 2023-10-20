package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorpreturngoodsrefuseAPIRequest 卖家拒绝退货 API请求
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
type TaobaorpreturngoodsrefuseAPIRequest struct {
	model.Params
	// 退款服务状态，售后或者售中
	_refundPhase string
	// 退款编号
	_refundId int64
	// 退款版本号
	_refundVersion int64
	// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
	_refuseProof *model.File
	// 拒绝原因编号，会提供拒绝原因列表供选择
	_refuseReasonId int64
}

// NewTaobaorpreturngoodsrefuseRequest 初始化TaobaorpreturngoodsrefuseAPIRequest对象
func NewTaobaorpreturngoodsrefuseRequest() *TaobaorpreturngoodsrefuseAPIRequest {
	return &TaobaorpreturngoodsrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorpreturngoodsrefuseAPIRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorpreturngoodsrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundPhase is RefundPhase Setter
// 退款服务状态，售后或者售中
func (r *TaobaorpreturngoodsrefuseAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorpreturngoodsrefuseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaorpreturngoodsrefuseAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetRefuseProof is RefuseProof Setter
// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
func (r *TaobaorpreturngoodsrefuseAPIRequest) SetRefuseProof(_refuseProof *model.File) error {
	r._refuseProof = _refuseProof
	r.Set("refuse_proof", _refuseProof)
	return nil
}

// GetRefuseProof RefuseProof Getter
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRefuseProof() *model.File {
	return r._refuseProof
}

// SetRefuseReasonId is RefuseReasonId Setter
// 拒绝原因编号，会提供拒绝原因列表供选择
func (r *TaobaorpreturngoodsrefuseAPIRequest) SetRefuseReasonId(_refuseReasonId int64) error {
	r._refuseReasonId = _refuseReasonId
	r.Set("refuse_reason_id", _refuseReasonId)
	return nil
}

// GetRefuseReasonId RefuseReasonId Getter
func (r TaobaorpreturngoodsrefuseAPIRequest) GetRefuseReasonId() int64 {
	return r._refuseReasonId
}