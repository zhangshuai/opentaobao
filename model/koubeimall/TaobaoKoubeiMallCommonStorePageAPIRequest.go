package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstorepageAPIRequest 分页查询综合体内的门店列表信息 API请求
// taobao.koubei.mall.common.store.page
//
// 分页查询综合体内的门店列表信息
type TaobaokoubeimallcommonstorepageAPIRequest struct {
	model.Params
	// 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
	_serviceTag []string
	// 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
	_categoryIds []string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 纬度（终端设备地理位置）
	_latitude string
	// 经度（终端设备地理位置）
	_longitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
	_order string
	// 商圈内的门店ID
	_storeId string
	// 商圈ID
	_mallId string
	// 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
	_pageSize int64
	// 分页查询起始值，默认为0
	_start int64
}

// NewTaobaokoubeimallcommonstorepageRequest 初始化TaobaokoubeimallcommonstorepageAPIRequest对象
func NewTaobaokoubeimallcommonstorepageRequest() *TaobaokoubeimallcommonstorepageAPIRequest {
	return &TaobaokoubeimallcommonstorepageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.store.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceTag is ServiceTag Setter
// 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetServiceTag(_serviceTag []string) error {
	r._serviceTag = _serviceTag
	r.Set("service_tag", _serviceTag)
	return nil
}

// GetServiceTag ServiceTag Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetServiceTag() []string {
	return r._serviceTag
}

// SetCategoryIds is CategoryIds Setter
// 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetCategoryIds(_categoryIds []string) error {
	r._categoryIds = _categoryIds
	r.Set("category_ids", _categoryIds)
	return nil
}

// GetCategoryIds CategoryIds Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetCategoryIds() []string {
	return r._categoryIds
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetOrder is Order Setter
// 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetOrder(_order string) error {
	r._order = _order
	r.Set("order", _order)
	return nil
}

// GetOrder Order Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetOrder() string {
	return r._order
}

// SetStoreId is StoreId Setter
// 商圈内的门店ID
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetMallId() string {
	return r._mallId
}

// SetPageSize is PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStart is Start Setter
// 分页查询起始值，默认为0
func (r *TaobaokoubeimallcommonstorepageAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaokoubeimallcommonstorepageAPIRequest) GetStart() int64 {
	return r._start
}
