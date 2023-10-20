package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchantenterpriseentryAPIRequest 机构入驻 API请求
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
type AlibabaalihousemerchantenterpriseentryAPIRequest struct {
	model.Params
	// KA名称
	_name string
	// 旗舰店nick
	_nick string
	// 外部旗舰店ID
	_outerId string
	// 店铺电话
	_storePhone string
	// 是否测试 0-否 1-是
	_isTest *model.File
	// 店铺类型：1-渠道店 2-官店
	_storeType int64
}

// NewAlibabaalihousemerchantenterpriseentryRequest 初始化AlibabaalihousemerchantenterpriseentryAPIRequest对象
func NewAlibabaalihousemerchantenterpriseentryRequest() *AlibabaalihousemerchantenterpriseentryAPIRequest {
	return &AlibabaalihousemerchantenterpriseentryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.merchant.enterprise.entry"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// KA名称
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetName() string {
	return r._name
}

// SetNick is Nick Setter
// 旗舰店nick
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetNick() string {
	return r._nick
}

// SetOuterId is OuterId Setter
// 外部旗舰店ID
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetStorePhone is StorePhone Setter
// 店铺电话
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetStorePhone(_storePhone string) error {
	r._storePhone = _storePhone
	r.Set("store_phone", _storePhone)
	return nil
}

// GetStorePhone StorePhone Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetStorePhone() string {
	return r._storePhone
}

// SetIsTest is IsTest Setter
// 是否测试 0-否 1-是
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetIsTest(_isTest *model.File) error {
	r._isTest = _isTest
	r.Set("is_test", _isTest)
	return nil
}

// GetIsTest IsTest Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetIsTest() *model.File {
	return r._isTest
}

// SetStoreType is StoreType Setter
// 店铺类型：1-渠道店 2-官店
func (r *AlibabaalihousemerchantenterpriseentryAPIRequest) SetStoreType(_storeType int64) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r AlibabaalihousemerchantenterpriseentryAPIRequest) GetStoreType() int64 {
	return r._storeType
}
