package aecreatives

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateHotproductDownloadAPIRequest 联盟营销爆品下载接口 API请求
// aliexpress.affiliate.hotproduct.download
//
// 查询联盟爆品API
type AliexpressAffiliateHotproductDownloadAPIRequest struct {
	model.Params
	// trackingId
	_trackingId string
	// 请求签名
	_appSignature string
	// 类目ID
	_categoryId string
	// 返回字段列表
	_fields string
	// 站点商品标：global,it_site,es_site,ru_site
	_localeSite string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
	// 请求页数
	_pageNo int64
	// 每次请求数量
	_pageSize int64
}

// NewAliexpressAffiliateHotproductDownloadRequest 初始化AliexpressAffiliateHotproductDownloadAPIRequest对象
func NewAliexpressAffiliateHotproductDownloadRequest() *AliexpressAffiliateHotproductDownloadAPIRequest {
	return &AliexpressAffiliateHotproductDownloadAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) Reset() {
	r._trackingId = ""
	r._appSignature = ""
	r._categoryId = ""
	r._fields = ""
	r._localeSite = ""
	r._targetCurrency = ""
	r._targetLanguage = ""
	r._country = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.hotproduct.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackingId is TrackingId Setter
// trackingId
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetAppSignature is AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetCategoryId is CategoryId Setter
// 类目ID
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetCategoryId(_categoryId string) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetCategoryId() string {
	return r._categoryId
}

// SetFields is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetFields() string {
	return r._fields
}

// SetLocaleSite is LocaleSite Setter
// 站点商品标：global,it_site,es_site,ru_site
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetLocaleSite(_localeSite string) error {
	r._localeSite = _localeSite
	r.Set("locale_site", _localeSite)
	return nil
}

// GetLocaleSite LocaleSite Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetLocaleSite() string {
	return r._localeSite
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetCountry is Country Setter
// 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetCountry() string {
	return r._country
}

// SetPageNo is PageNo Setter
// 请求页数
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每次请求数量
func (r *AliexpressAffiliateHotproductDownloadAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressAffiliateHotproductDownloadAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAliexpressAffiliateHotproductDownloadAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAffiliateHotproductDownloadRequest()
	},
}

// GetAliexpressAffiliateHotproductDownloadRequest 从 sync.Pool 获取 AliexpressAffiliateHotproductDownloadAPIRequest
func GetAliexpressAffiliateHotproductDownloadAPIRequest() *AliexpressAffiliateHotproductDownloadAPIRequest {
	return poolAliexpressAffiliateHotproductDownloadAPIRequest.Get().(*AliexpressAffiliateHotproductDownloadAPIRequest)
}

// ReleaseAliexpressAffiliateHotproductDownloadAPIRequest 将 AliexpressAffiliateHotproductDownloadAPIRequest 放入 sync.Pool
func ReleaseAliexpressAffiliateHotproductDownloadAPIRequest(v *AliexpressAffiliateHotproductDownloadAPIRequest) {
	v.Reset()
	poolAliexpressAffiliateHotproductDownloadAPIRequest.Put(v)
}
