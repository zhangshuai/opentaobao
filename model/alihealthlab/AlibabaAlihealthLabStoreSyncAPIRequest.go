package alihealthlab

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabStoreSyncAPIRequest 阿里健康检验检测业务，isv门店同步到健康 API请求
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
type AlibabaAlihealthLabStoreSyncAPIRequest struct {
	model.Params
	// 支持的淘宝商品类目ID，阿里医院场景
	_allowedTbItemCategoryIds []int64
	// EFFECTIVE 生效，INVALID 失效
	_isvStoreStatus string
	// 预约须知
	_reserveNotice string
	// 门店类型描述
	_storeTypeDesc string
	// 企业社会征信号
	_socialCreditCode string
	// 营业执照编号
	_licenseNo string
	// 营业执照名称
	_licenseName string
	// 门店交通路线
	_storeRoutesDesc string
	// 营业时间描述
	_workTimeDesc string
	// 门店电话号码
	_storePhone string
	// 门店介绍
	_storeIntro string
	// 门店地址
	_storeAddress string
	// isv门店编码
	_isvStoreCode string
	// 门店名称
	_storeName string
	// 经度
	_longitude float64
	// 纬度
	_latitude float64
	// 城市编码
	_cityCode int64
	// 支持在线报告
	_supportOnlineReport bool
}

// NewAlibabaAlihealthLabStoreSyncRequest 初始化AlibabaAlihealthLabStoreSyncAPIRequest对象
func NewAlibabaAlihealthLabStoreSyncRequest() *AlibabaAlihealthLabStoreSyncAPIRequest {
	return &AlibabaAlihealthLabStoreSyncAPIRequest{
		Params: model.NewParams(18),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) Reset() {
	r._allowedTbItemCategoryIds = r._allowedTbItemCategoryIds[:0]
	r._isvStoreStatus = ""
	r._reserveNotice = ""
	r._storeTypeDesc = ""
	r._socialCreditCode = ""
	r._licenseNo = ""
	r._licenseName = ""
	r._storeRoutesDesc = ""
	r._workTimeDesc = ""
	r._storePhone = ""
	r._storeIntro = ""
	r._storeAddress = ""
	r._isvStoreCode = ""
	r._storeName = ""
	r._longitude = 0
	r._latitude = 0
	r._cityCode = 0
	r._supportOnlineReport = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.store.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAllowedTbItemCategoryIds is AllowedTbItemCategoryIds Setter
// 支持的淘宝商品类目ID，阿里医院场景
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetAllowedTbItemCategoryIds(_allowedTbItemCategoryIds []int64) error {
	r._allowedTbItemCategoryIds = _allowedTbItemCategoryIds
	r.Set("allowed_tb_item_category_ids", _allowedTbItemCategoryIds)
	return nil
}

// GetAllowedTbItemCategoryIds AllowedTbItemCategoryIds Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetAllowedTbItemCategoryIds() []int64 {
	return r._allowedTbItemCategoryIds
}

// SetIsvStoreStatus is IsvStoreStatus Setter
// EFFECTIVE 生效，INVALID 失效
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetIsvStoreStatus(_isvStoreStatus string) error {
	r._isvStoreStatus = _isvStoreStatus
	r.Set("isv_store_status", _isvStoreStatus)
	return nil
}

// GetIsvStoreStatus IsvStoreStatus Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetIsvStoreStatus() string {
	return r._isvStoreStatus
}

// SetReserveNotice is ReserveNotice Setter
// 预约须知
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetReserveNotice(_reserveNotice string) error {
	r._reserveNotice = _reserveNotice
	r.Set("reserve_notice", _reserveNotice)
	return nil
}

// GetReserveNotice ReserveNotice Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetReserveNotice() string {
	return r._reserveNotice
}

// SetStoreTypeDesc is StoreTypeDesc Setter
// 门店类型描述
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStoreTypeDesc(_storeTypeDesc string) error {
	r._storeTypeDesc = _storeTypeDesc
	r.Set("store_type_desc", _storeTypeDesc)
	return nil
}

// GetStoreTypeDesc StoreTypeDesc Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStoreTypeDesc() string {
	return r._storeTypeDesc
}

// SetSocialCreditCode is SocialCreditCode Setter
// 企业社会征信号
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetSocialCreditCode(_socialCreditCode string) error {
	r._socialCreditCode = _socialCreditCode
	r.Set("social_credit_code", _socialCreditCode)
	return nil
}

// GetSocialCreditCode SocialCreditCode Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetSocialCreditCode() string {
	return r._socialCreditCode
}

// SetLicenseNo is LicenseNo Setter
// 营业执照编号
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetLicenseNo(_licenseNo string) error {
	r._licenseNo = _licenseNo
	r.Set("license_no", _licenseNo)
	return nil
}

// GetLicenseNo LicenseNo Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetLicenseNo() string {
	return r._licenseNo
}

// SetLicenseName is LicenseName Setter
// 营业执照名称
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetLicenseName(_licenseName string) error {
	r._licenseName = _licenseName
	r.Set("license_name", _licenseName)
	return nil
}

// GetLicenseName LicenseName Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetLicenseName() string {
	return r._licenseName
}

// SetStoreRoutesDesc is StoreRoutesDesc Setter
// 门店交通路线
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStoreRoutesDesc(_storeRoutesDesc string) error {
	r._storeRoutesDesc = _storeRoutesDesc
	r.Set("store_routes_desc", _storeRoutesDesc)
	return nil
}

// GetStoreRoutesDesc StoreRoutesDesc Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStoreRoutesDesc() string {
	return r._storeRoutesDesc
}

// SetWorkTimeDesc is WorkTimeDesc Setter
// 营业时间描述
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetWorkTimeDesc(_workTimeDesc string) error {
	r._workTimeDesc = _workTimeDesc
	r.Set("work_time_desc", _workTimeDesc)
	return nil
}

// GetWorkTimeDesc WorkTimeDesc Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetWorkTimeDesc() string {
	return r._workTimeDesc
}

// SetStorePhone is StorePhone Setter
// 门店电话号码
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStorePhone(_storePhone string) error {
	r._storePhone = _storePhone
	r.Set("store_phone", _storePhone)
	return nil
}

// GetStorePhone StorePhone Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStorePhone() string {
	return r._storePhone
}

// SetStoreIntro is StoreIntro Setter
// 门店介绍
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStoreIntro(_storeIntro string) error {
	r._storeIntro = _storeIntro
	r.Set("store_intro", _storeIntro)
	return nil
}

// GetStoreIntro StoreIntro Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStoreIntro() string {
	return r._storeIntro
}

// SetStoreAddress is StoreAddress Setter
// 门店地址
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStoreAddress(_storeAddress string) error {
	r._storeAddress = _storeAddress
	r.Set("store_address", _storeAddress)
	return nil
}

// GetStoreAddress StoreAddress Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStoreAddress() string {
	return r._storeAddress
}

// SetIsvStoreCode is IsvStoreCode Setter
// isv门店编码
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetIsvStoreCode(_isvStoreCode string) error {
	r._isvStoreCode = _isvStoreCode
	r.Set("isv_store_code", _isvStoreCode)
	return nil
}

// GetIsvStoreCode IsvStoreCode Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetIsvStoreCode() string {
	return r._isvStoreCode
}

// SetStoreName is StoreName Setter
// 门店名称
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetLongitude is Longitude Setter
// 经度
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetLongitude(_longitude float64) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetLongitude() float64 {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetLatitude(_latitude float64) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetLatitude() float64 {
	return r._latitude
}

// SetCityCode is CityCode Setter
// 城市编码
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetSupportOnlineReport is SupportOnlineReport Setter
// 支持在线报告
func (r *AlibabaAlihealthLabStoreSyncAPIRequest) SetSupportOnlineReport(_supportOnlineReport bool) error {
	r._supportOnlineReport = _supportOnlineReport
	r.Set("support_online_report", _supportOnlineReport)
	return nil
}

// GetSupportOnlineReport SupportOnlineReport Getter
func (r AlibabaAlihealthLabStoreSyncAPIRequest) GetSupportOnlineReport() bool {
	return r._supportOnlineReport
}

var poolAlibabaAlihealthLabStoreSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthLabStoreSyncRequest()
	},
}

// GetAlibabaAlihealthLabStoreSyncRequest 从 sync.Pool 获取 AlibabaAlihealthLabStoreSyncAPIRequest
func GetAlibabaAlihealthLabStoreSyncAPIRequest() *AlibabaAlihealthLabStoreSyncAPIRequest {
	return poolAlibabaAlihealthLabStoreSyncAPIRequest.Get().(*AlibabaAlihealthLabStoreSyncAPIRequest)
}

// ReleaseAlibabaAlihealthLabStoreSyncAPIRequest 将 AlibabaAlihealthLabStoreSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthLabStoreSyncAPIRequest(v *AlibabaAlihealthLabStoreSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthLabStoreSyncAPIRequest.Put(v)
}
