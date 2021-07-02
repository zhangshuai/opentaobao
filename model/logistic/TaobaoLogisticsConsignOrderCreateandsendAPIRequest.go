package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsConsignOrderCreateandsendAPIRequest 创建订单并发货 API请求
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendAPIRequest struct {
	model.Params
	// 用户ID
	_userId int64
	// 订单来源，值选择：30
	_orderSource int64
	// 订单类型，值固定选择：30
	_orderType int64
	// 物流订单物流类型，值固定选择：2
	_logisType int64
	// 物流公司ID
	_companyId int64
	// 交易流水号，淘外订单号或者商家内部交易流水号
	_tradeId int64
	// 运单号
	_mailNo string
	// 费用承担方式 1买家承担运费 2卖家承担运费
	_shipping string
	// 发件人名称
	_sName string
	// 发件人区域ID
	_sAreaId int64
	// 发件人街道地址
	_sAddress string
	// 发件人出编
	_sZipCode string
	// 手机号码
	_sMobilePhone string
	// 电话号码
	_sTelephone string
	// 省
	_sProvName string
	// 市
	_sCityName string
	// 区
	_sDistName string
	// 收件人名称
	_rName string
	// 收件人区域ID
	_rAreaId int64
	// 收件人街道地址
	_rAddress string
	// 收件人邮编
	_rZipCode string
	// 手机号码
	_rMobilePhone string
	// 电话号码
	_rTelephone string
	// 省
	_rProvName string
	// 市
	_rCityName string
	// 区
	_rDistName string
	// 物品的json数据。
	_itemJsonString string
}

// NewTaobaoLogisticsConsignOrderCreateandsendRequest 初始化TaobaoLogisticsConsignOrderCreateandsendAPIRequest对象
func NewTaobaoLogisticsConsignOrderCreateandsendRequest() *TaobaoLogisticsConsignOrderCreateandsendAPIRequest {
	return &TaobaoLogisticsConsignOrderCreateandsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.consign.order.createandsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户ID
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is OrderSource Setter
// 订单来源，值选择：30
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetOrderSource(_orderSource int64) error {
	r._orderSource = _orderSource
	r.Set("order_source", _orderSource)
	return nil
}

// Get OrderSource Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetOrderSource() int64 {
	return r._orderSource
}

// Set is OrderType Setter
// 订单类型，值固定选择：30
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// Get OrderType Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// Set is LogisType Setter
// 物流订单物流类型，值固定选择：2
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetLogisType(_logisType int64) error {
	r._logisType = _logisType
	r.Set("logis_type", _logisType)
	return nil
}

// Get LogisType Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetLogisType() int64 {
	return r._logisType
}

// Set is CompanyId Setter
// 物流公司ID
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// Get CompanyId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// Set is TradeId Setter
// 交易流水号，淘外订单号或者商家内部交易流水号
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// Get TradeId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// Set is MailNo Setter
// 运单号
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// Get MailNo Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetMailNo() string {
	return r._mailNo
}

// Set is Shipping Setter
// 费用承担方式 1买家承担运费 2卖家承担运费
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetShipping(_shipping string) error {
	r._shipping = _shipping
	r.Set("shipping", _shipping)
	return nil
}

// Get Shipping Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetShipping() string {
	return r._shipping
}

// Set is SName Setter
// 发件人名称
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSName(_sName string) error {
	r._sName = _sName
	r.Set("s_name", _sName)
	return nil
}

// Get SName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSName() string {
	return r._sName
}

// Set is SAreaId Setter
// 发件人区域ID
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSAreaId(_sAreaId int64) error {
	r._sAreaId = _sAreaId
	r.Set("s_area_id", _sAreaId)
	return nil
}

// Get SAreaId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSAreaId() int64 {
	return r._sAreaId
}

// Set is SAddress Setter
// 发件人街道地址
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSAddress(_sAddress string) error {
	r._sAddress = _sAddress
	r.Set("s_address", _sAddress)
	return nil
}

// Get SAddress Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSAddress() string {
	return r._sAddress
}

// Set is SZipCode Setter
// 发件人出编
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSZipCode(_sZipCode string) error {
	r._sZipCode = _sZipCode
	r.Set("s_zip_code", _sZipCode)
	return nil
}

// Get SZipCode Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSZipCode() string {
	return r._sZipCode
}

// Set is SMobilePhone Setter
// 手机号码
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSMobilePhone(_sMobilePhone string) error {
	r._sMobilePhone = _sMobilePhone
	r.Set("s_mobile_phone", _sMobilePhone)
	return nil
}

// Get SMobilePhone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSMobilePhone() string {
	return r._sMobilePhone
}

// Set is STelephone Setter
// 电话号码
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSTelephone(_sTelephone string) error {
	r._sTelephone = _sTelephone
	r.Set("s_telephone", _sTelephone)
	return nil
}

// Get STelephone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSTelephone() string {
	return r._sTelephone
}

// Set is SProvName Setter
// 省
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSProvName(_sProvName string) error {
	r._sProvName = _sProvName
	r.Set("s_prov_name", _sProvName)
	return nil
}

// Get SProvName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSProvName() string {
	return r._sProvName
}

// Set is SCityName Setter
// 市
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSCityName(_sCityName string) error {
	r._sCityName = _sCityName
	r.Set("s_city_name", _sCityName)
	return nil
}

// Get SCityName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSCityName() string {
	return r._sCityName
}

// Set is SDistName Setter
// 区
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetSDistName(_sDistName string) error {
	r._sDistName = _sDistName
	r.Set("s_dist_name", _sDistName)
	return nil
}

// Get SDistName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetSDistName() string {
	return r._sDistName
}

// Set is RName Setter
// 收件人名称
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRName(_rName string) error {
	r._rName = _rName
	r.Set("r_name", _rName)
	return nil
}

// Get RName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRName() string {
	return r._rName
}

// Set is RAreaId Setter
// 收件人区域ID
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRAreaId(_rAreaId int64) error {
	r._rAreaId = _rAreaId
	r.Set("r_area_id", _rAreaId)
	return nil
}

// Get RAreaId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRAreaId() int64 {
	return r._rAreaId
}

// Set is RAddress Setter
// 收件人街道地址
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRAddress(_rAddress string) error {
	r._rAddress = _rAddress
	r.Set("r_address", _rAddress)
	return nil
}

// Get RAddress Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRAddress() string {
	return r._rAddress
}

// Set is RZipCode Setter
// 收件人邮编
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRZipCode(_rZipCode string) error {
	r._rZipCode = _rZipCode
	r.Set("r_zip_code", _rZipCode)
	return nil
}

// Get RZipCode Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRZipCode() string {
	return r._rZipCode
}

// Set is RMobilePhone Setter
// 手机号码
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRMobilePhone(_rMobilePhone string) error {
	r._rMobilePhone = _rMobilePhone
	r.Set("r_mobile_phone", _rMobilePhone)
	return nil
}

// Get RMobilePhone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRMobilePhone() string {
	return r._rMobilePhone
}

// Set is RTelephone Setter
// 电话号码
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRTelephone(_rTelephone string) error {
	r._rTelephone = _rTelephone
	r.Set("r_telephone", _rTelephone)
	return nil
}

// Get RTelephone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRTelephone() string {
	return r._rTelephone
}

// Set is RProvName Setter
// 省
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRProvName(_rProvName string) error {
	r._rProvName = _rProvName
	r.Set("r_prov_name", _rProvName)
	return nil
}

// Get RProvName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRProvName() string {
	return r._rProvName
}

// Set is RCityName Setter
// 市
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRCityName(_rCityName string) error {
	r._rCityName = _rCityName
	r.Set("r_city_name", _rCityName)
	return nil
}

// Get RCityName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRCityName() string {
	return r._rCityName
}

// Set is RDistName Setter
// 区
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetRDistName(_rDistName string) error {
	r._rDistName = _rDistName
	r.Set("r_dist_name", _rDistName)
	return nil
}

// Get RDistName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetRDistName() string {
	return r._rDistName
}

// Set is ItemJsonString Setter
// 物品的json数据。
func (r *TaobaoLogisticsConsignOrderCreateandsendAPIRequest) SetItemJsonString(_itemJsonString string) error {
	r._itemJsonString = _itemJsonString
	r.Set("item_json_string", _itemJsonString)
	return nil
}

// Get ItemJsonString Getter
func (r TaobaoLogisticsConsignOrderCreateandsendAPIRequest) GetItemJsonString() string {
	return r._itemJsonString
}
