package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest 查询门店推荐菜信息 API请求
// taobao.koubei.mall.common.store.display.goods.list
//
// 提供查询口碑商圈内的门店推荐菜信息
type TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest struct {
	model.Params
	// 商圈ID
	_mallId string
	// 门店ID
	_storeId string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
}

// NewTaobaokoubeimallcommonstoredisplaygoodslistRequest 初始化TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest对象
func NewTaobaokoubeimallcommonstoredisplaygoodslistRequest() *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest {
	return &TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.store.display.goods.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetMallId() string {
	return r._mallId
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonstoredisplaygoodslistAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}
