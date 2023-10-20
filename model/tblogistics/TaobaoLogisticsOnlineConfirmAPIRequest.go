package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsonlineconfirmAPIRequest 确认发货通知接口 API请求
// taobao.logistics.online.confirm
//
// &lt;br&gt;&lt;font color=&#39;red&#39;&gt;仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。&lt;br&gt;
// 确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。&lt;/font&gt;
type TaobaologisticsonlineconfirmAPIRequest struct {
	model.Params
	// 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
	_subTid []int64
	// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
	_outSid string
	// 商家的IP地址
	_sellerIp string
	// 淘宝交易ID
	_tid int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
}

// NewTaobaologisticsonlineconfirmRequest 初始化TaobaologisticsonlineconfirmAPIRequest对象
func NewTaobaologisticsonlineconfirmRequest() *TaobaologisticsonlineconfirmAPIRequest {
	return &TaobaologisticsonlineconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsonlineconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.online.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsonlineconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsonlineconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubTid is SubTid Setter
// 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
func (r *TaobaologisticsonlineconfirmAPIRequest) SetSubTid(_subTid []int64) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r TaobaologisticsonlineconfirmAPIRequest) GetSubTid() []int64 {
	return r._subTid
}

// SetOutSid is OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaologisticsonlineconfirmAPIRequest) SetOutSid(_outSid string) error {
	r._outSid = _outSid
	r.Set("out_sid", _outSid)
	return nil
}

// GetOutSid OutSid Getter
func (r TaobaologisticsonlineconfirmAPIRequest) GetOutSid() string {
	return r._outSid
}

// SetSellerIp is SellerIp Setter
// 商家的IP地址
func (r *TaobaologisticsonlineconfirmAPIRequest) SetSellerIp(_sellerIp string) error {
	r._sellerIp = _sellerIp
	r.Set("seller_ip", _sellerIp)
	return nil
}

// GetSellerIp SellerIp Getter
func (r TaobaologisticsonlineconfirmAPIRequest) GetSellerIp() string {
	return r._sellerIp
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaologisticsonlineconfirmAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticsonlineconfirmAPIRequest) GetTid() int64 {
	return r._tid
}

// SetIsSplit is IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaologisticsonlineconfirmAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// GetIsSplit IsSplit Getter
func (r TaobaologisticsonlineconfirmAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}