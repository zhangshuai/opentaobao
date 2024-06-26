package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkCouponConvertAPIRequest 淘宝客-推广者-单品券高效转链 API请求
// taobao.tbk.coupon.convert
//
// 单品券高效转链API
type TaobaoTbkCouponConvertAPIRequest struct {
	model.Params
	// 淘客商品id
	_itemId string
	// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
	_relationId string
	// 会员运营ID
	_specialId string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalId string
	// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
	_xid string
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
	_bizSceneId string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，其它模式不用填写本字段）
	_promotionType string
	// 推广位id，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
	// 1：PC，2：无线，默认：１
	_platform int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdId int64
	// 是否获取前N件佣金	,0-否，1-是,其他值-否
	_getTopnRate int64
	// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
	_miniProgramLink int64
	// 商品库服务账户(场景id3权限对应的memberid）
	_manageItemPubId int64
	// 入参商品id下的skuid，传入时会透传至转链结果url中
	_skuId int64
}

// NewTaobaoTbkCouponConvertRequest 初始化TaobaoTbkCouponConvertAPIRequest对象
func NewTaobaoTbkCouponConvertRequest() *TaobaoTbkCouponConvertAPIRequest {
	return &TaobaoTbkCouponConvertAPIRequest{
		Params: model.NewParams(14),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkCouponConvertAPIRequest) Reset() {
	r._itemId = ""
	r._relationId = ""
	r._specialId = ""
	r._externalId = ""
	r._xid = ""
	r._bizSceneId = ""
	r._promotionType = ""
	r._adzoneId = 0
	r._platform = 0
	r._ucrowdId = 0
	r._getTopnRate = 0
	r._miniProgramLink = 0
	r._manageItemPubId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkCouponConvertAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.coupon.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkCouponConvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkCouponConvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 淘客商品id
func (r *TaobaoTbkCouponConvertAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetItemId() string {
	return r._itemId
}

// SetRelationId is RelationId Setter
// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
func (r *TaobaoTbkCouponConvertAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSpecialId is SpecialId Setter
// 会员运营ID
func (r *TaobaoTbkCouponConvertAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetExternalId is ExternalId Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaoTbkCouponConvertAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetExternalId() string {
	return r._externalId
}

// SetXid is Xid Setter
// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
func (r *TaobaoTbkCouponConvertAPIRequest) SetXid(_xid string) error {
	r._xid = _xid
	r.Set("xid", _xid)
	return nil
}

// GetXid Xid Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetXid() string {
	return r._xid
}

// SetBizSceneId is BizSceneId Setter
// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
func (r *TaobaoTbkCouponConvertAPIRequest) SetBizSceneId(_bizSceneId string) error {
	r._bizSceneId = _bizSceneId
	r.Set("biz_scene_id", _bizSceneId)
	return nil
}

// GetBizSceneId BizSceneId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetBizSceneId() string {
	return r._bizSceneId
}

// SetPromotionType is PromotionType Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，其它模式不用填写本字段）
func (r *TaobaoTbkCouponConvertAPIRequest) SetPromotionType(_promotionType string) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetPromotionType() string {
	return r._promotionType
}

// SetAdzoneId is AdzoneId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkCouponConvertAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetPlatform is Platform Setter
// 1：PC，2：无线，默认：１
func (r *TaobaoTbkCouponConvertAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetUcrowdId is UcrowdId Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaoTbkCouponConvertAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetGetTopnRate is GetTopnRate Setter
// 是否获取前N件佣金	,0-否，1-是,其他值-否
func (r *TaobaoTbkCouponConvertAPIRequest) SetGetTopnRate(_getTopnRate int64) error {
	r._getTopnRate = _getTopnRate
	r.Set("get_topn_rate", _getTopnRate)
	return nil
}

// GetGetTopnRate GetTopnRate Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetGetTopnRate() int64 {
	return r._getTopnRate
}

// SetMiniProgramLink is MiniProgramLink Setter
// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
func (r *TaobaoTbkCouponConvertAPIRequest) SetMiniProgramLink(_miniProgramLink int64) error {
	r._miniProgramLink = _miniProgramLink
	r.Set("mini_program_link", _miniProgramLink)
	return nil
}

// GetMiniProgramLink MiniProgramLink Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetMiniProgramLink() int64 {
	return r._miniProgramLink
}

// SetManageItemPubId is ManageItemPubId Setter
// 商品库服务账户(场景id3权限对应的memberid）
func (r *TaobaoTbkCouponConvertAPIRequest) SetManageItemPubId(_manageItemPubId int64) error {
	r._manageItemPubId = _manageItemPubId
	r.Set("manage_item_pub_id", _manageItemPubId)
	return nil
}

// GetManageItemPubId ManageItemPubId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetManageItemPubId() int64 {
	return r._manageItemPubId
}

// SetSkuId is SkuId Setter
// 入参商品id下的skuid，传入时会透传至转链结果url中
func (r *TaobaoTbkCouponConvertAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoTbkCouponConvertAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoTbkCouponConvertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkCouponConvertRequest()
	},
}

// GetTaobaoTbkCouponConvertRequest 从 sync.Pool 获取 TaobaoTbkCouponConvertAPIRequest
func GetTaobaoTbkCouponConvertAPIRequest() *TaobaoTbkCouponConvertAPIRequest {
	return poolTaobaoTbkCouponConvertAPIRequest.Get().(*TaobaoTbkCouponConvertAPIRequest)
}

// ReleaseTaobaoTbkCouponConvertAPIRequest 将 TaobaoTbkCouponConvertAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkCouponConvertAPIRequest(v *TaobaoTbkCouponConvertAPIRequest) {
	v.Reset()
	poolTaobaoTbkCouponConvertAPIRequest.Put(v)
}
