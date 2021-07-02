package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoPutAPIRequest 电子价签显示用商品信息写入 API请求
// taobao.uscesl.iteminfo.put
//
// 用于电子价签上显示的商品信息的写入，包含价格及促销信息
type TaobaoUsceslIteminfoPutAPIRequest struct {
	model.Params
	// 型号
	_modelNum string
	// 价格单位
	_priceUnit string
	// 品牌名
	_brandName string
	// 销售规格
	_saleSpec string
	// 分类
	_categoryName string
	// 等级
	_rank string
	// 商品变更状态
	_itemChangeStatus string
	// 实际销售价格，单位：分
	_acctionPrice string
	// 能效
	_energyEfficiency string
	// 商品skuId
	_skuId string
	// 促销开始时间:yyyy-MM-dd HH:mm:ss
	_promotionStart string
	// 商品条码
	_itemBarCode string
	// 商品名称
	_itemTitle string
	// 促销文案
	_promotionText string
	// 扩展属性C
	_customizeFeatureC string
	// 扩展属性D
	_customizeFeatureD string
	// 扩展属性E
	_customizeFeatureE string
	// 扩展属性F
	_customizeFeatureF string
	// 扩展属性G
	_customizeFeatureG string
	// 扩展属性H
	_customizeFeatureH string
	// 扩展属性I
	_customizeFeatureI string
	// 扩展属性J
	_customizeFeatureJ string
	// 二维码
	_itemQrCode string
	// 商品状态0:在售 1:售罄
	_itemStatus int64
	// 促销状态0:非促销 1:促销
	_ifPromotion bool
	// 商品编码
	_itemId int64
	// 促销结束时间:yyyy-MM-dd HH:mm:ss
	_promotionEnd string
	// 促销原因
	_promotionReason string
	// 商品原价
	_originalPrice string
	// 商品简称
	_shortTitle string
	// 扩展属性B
	_customizeFeatureB string
	// 产地
	_productionPlace string
	// 扩展属性A
	_customizeFeatureA string
	// 门店ID
	_shopId int64
}

// NewTaobaoUsceslIteminfoPutRequest 初始化TaobaoUsceslIteminfoPutAPIRequest对象
func NewTaobaoUsceslIteminfoPutRequest() *TaobaoUsceslIteminfoPutAPIRequest {
	return &TaobaoUsceslIteminfoPutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoPutAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.iteminfo.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoPutAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ModelNum Setter
// 型号
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetModelNum(_modelNum string) error {
	r._modelNum = _modelNum
	r.Set("model_num", _modelNum)
	return nil
}

// Get ModelNum Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetModelNum() string {
	return r._modelNum
}

// Set is PriceUnit Setter
// 价格单位
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetPriceUnit(_priceUnit string) error {
	r._priceUnit = _priceUnit
	r.Set("price_unit", _priceUnit)
	return nil
}

// Get PriceUnit Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetPriceUnit() string {
	return r._priceUnit
}

// Set is BrandName Setter
// 品牌名
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// Get BrandName Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetBrandName() string {
	return r._brandName
}

// Set is SaleSpec Setter
// 销售规格
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetSaleSpec(_saleSpec string) error {
	r._saleSpec = _saleSpec
	r.Set("sale_spec", _saleSpec)
	return nil
}

// Get SaleSpec Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetSaleSpec() string {
	return r._saleSpec
}

// Set is CategoryName Setter
// 分类
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// Get CategoryName Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// Set is Rank Setter
// 等级
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetRank(_rank string) error {
	r._rank = _rank
	r.Set("rank", _rank)
	return nil
}

// Get Rank Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetRank() string {
	return r._rank
}

// Set is ItemChangeStatus Setter
// 商品变更状态
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemChangeStatus(_itemChangeStatus string) error {
	r._itemChangeStatus = _itemChangeStatus
	r.Set("item_change_status", _itemChangeStatus)
	return nil
}

// Get ItemChangeStatus Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemChangeStatus() string {
	return r._itemChangeStatus
}

// Set is AcctionPrice Setter
// 实际销售价格，单位：分
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetAcctionPrice(_acctionPrice string) error {
	r._acctionPrice = _acctionPrice
	r.Set("acction_price", _acctionPrice)
	return nil
}

// Get AcctionPrice Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetAcctionPrice() string {
	return r._acctionPrice
}

// Set is EnergyEfficiency Setter
// 能效
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetEnergyEfficiency(_energyEfficiency string) error {
	r._energyEfficiency = _energyEfficiency
	r.Set("energy_efficiency", _energyEfficiency)
	return nil
}

// Get EnergyEfficiency Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetEnergyEfficiency() string {
	return r._energyEfficiency
}

// Set is SkuId Setter
// 商品skuId
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetSkuId(_skuId string) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetSkuId() string {
	return r._skuId
}

// Set is PromotionStart Setter
// 促销开始时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetPromotionStart(_promotionStart string) error {
	r._promotionStart = _promotionStart
	r.Set("promotion_start", _promotionStart)
	return nil
}

// Get PromotionStart Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetPromotionStart() string {
	return r._promotionStart
}

// Set is ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// Get ItemBarCode Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// Set is ItemTitle Setter
// 商品名称
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemTitle(_itemTitle string) error {
	r._itemTitle = _itemTitle
	r.Set("item_title", _itemTitle)
	return nil
}

// Get ItemTitle Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemTitle() string {
	return r._itemTitle
}

// Set is PromotionText Setter
// 促销文案
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetPromotionText(_promotionText string) error {
	r._promotionText = _promotionText
	r.Set("promotion_text", _promotionText)
	return nil
}

// Get PromotionText Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetPromotionText() string {
	return r._promotionText
}

// Set is CustomizeFeatureC Setter
// 扩展属性C
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureC(_customizeFeatureC string) error {
	r._customizeFeatureC = _customizeFeatureC
	r.Set("customize_feature_c", _customizeFeatureC)
	return nil
}

// Get CustomizeFeatureC Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureC() string {
	return r._customizeFeatureC
}

// Set is CustomizeFeatureD Setter
// 扩展属性D
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureD(_customizeFeatureD string) error {
	r._customizeFeatureD = _customizeFeatureD
	r.Set("customize_feature_d", _customizeFeatureD)
	return nil
}

// Get CustomizeFeatureD Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureD() string {
	return r._customizeFeatureD
}

// Set is CustomizeFeatureE Setter
// 扩展属性E
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureE(_customizeFeatureE string) error {
	r._customizeFeatureE = _customizeFeatureE
	r.Set("customize_feature_e", _customizeFeatureE)
	return nil
}

// Get CustomizeFeatureE Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureE() string {
	return r._customizeFeatureE
}

// Set is CustomizeFeatureF Setter
// 扩展属性F
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureF(_customizeFeatureF string) error {
	r._customizeFeatureF = _customizeFeatureF
	r.Set("customize_feature_f", _customizeFeatureF)
	return nil
}

// Get CustomizeFeatureF Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureF() string {
	return r._customizeFeatureF
}

// Set is CustomizeFeatureG Setter
// 扩展属性G
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureG(_customizeFeatureG string) error {
	r._customizeFeatureG = _customizeFeatureG
	r.Set("customize_feature_g", _customizeFeatureG)
	return nil
}

// Get CustomizeFeatureG Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureG() string {
	return r._customizeFeatureG
}

// Set is CustomizeFeatureH Setter
// 扩展属性H
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureH(_customizeFeatureH string) error {
	r._customizeFeatureH = _customizeFeatureH
	r.Set("customize_feature_h", _customizeFeatureH)
	return nil
}

// Get CustomizeFeatureH Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureH() string {
	return r._customizeFeatureH
}

// Set is CustomizeFeatureI Setter
// 扩展属性I
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureI(_customizeFeatureI string) error {
	r._customizeFeatureI = _customizeFeatureI
	r.Set("customize_feature_i", _customizeFeatureI)
	return nil
}

// Get CustomizeFeatureI Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureI() string {
	return r._customizeFeatureI
}

// Set is CustomizeFeatureJ Setter
// 扩展属性J
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureJ(_customizeFeatureJ string) error {
	r._customizeFeatureJ = _customizeFeatureJ
	r.Set("customize_feature_j", _customizeFeatureJ)
	return nil
}

// Get CustomizeFeatureJ Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureJ() string {
	return r._customizeFeatureJ
}

// Set is ItemQrCode Setter
// 二维码
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemQrCode(_itemQrCode string) error {
	r._itemQrCode = _itemQrCode
	r.Set("item_qr_code", _itemQrCode)
	return nil
}

// Get ItemQrCode Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemQrCode() string {
	return r._itemQrCode
}

// Set is ItemStatus Setter
// 商品状态0:在售 1:售罄
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemStatus(_itemStatus int64) error {
	r._itemStatus = _itemStatus
	r.Set("item_status", _itemStatus)
	return nil
}

// Get ItemStatus Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemStatus() int64 {
	return r._itemStatus
}

// Set is IfPromotion Setter
// 促销状态0:非促销 1:促销
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetIfPromotion(_ifPromotion bool) error {
	r._ifPromotion = _ifPromotion
	r.Set("if_promotion", _ifPromotion)
	return nil
}

// Get IfPromotion Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetIfPromotion() bool {
	return r._ifPromotion
}

// Set is ItemId Setter
// 商品编码
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is PromotionEnd Setter
// 促销结束时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetPromotionEnd(_promotionEnd string) error {
	r._promotionEnd = _promotionEnd
	r.Set("promotion_end", _promotionEnd)
	return nil
}

// Get PromotionEnd Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetPromotionEnd() string {
	return r._promotionEnd
}

// Set is PromotionReason Setter
// 促销原因
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetPromotionReason(_promotionReason string) error {
	r._promotionReason = _promotionReason
	r.Set("promotion_reason", _promotionReason)
	return nil
}

// Get PromotionReason Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetPromotionReason() string {
	return r._promotionReason
}

// Set is OriginalPrice Setter
// 商品原价
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetOriginalPrice(_originalPrice string) error {
	r._originalPrice = _originalPrice
	r.Set("original_price", _originalPrice)
	return nil
}

// Get OriginalPrice Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetOriginalPrice() string {
	return r._originalPrice
}

// Set is ShortTitle Setter
// 商品简称
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetShortTitle(_shortTitle string) error {
	r._shortTitle = _shortTitle
	r.Set("short_title", _shortTitle)
	return nil
}

// Get ShortTitle Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetShortTitle() string {
	return r._shortTitle
}

// Set is CustomizeFeatureB Setter
// 扩展属性B
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureB(_customizeFeatureB string) error {
	r._customizeFeatureB = _customizeFeatureB
	r.Set("customize_feature_b", _customizeFeatureB)
	return nil
}

// Get CustomizeFeatureB Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureB() string {
	return r._customizeFeatureB
}

// Set is ProductionPlace Setter
// 产地
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetProductionPlace(_productionPlace string) error {
	r._productionPlace = _productionPlace
	r.Set("production_place", _productionPlace)
	return nil
}

// Get ProductionPlace Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetProductionPlace() string {
	return r._productionPlace
}

// Set is CustomizeFeatureA Setter
// 扩展属性A
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetCustomizeFeatureA(_customizeFeatureA string) error {
	r._customizeFeatureA = _customizeFeatureA
	r.Set("customize_feature_a", _customizeFeatureA)
	return nil
}

// Get CustomizeFeatureA Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetCustomizeFeatureA() string {
	return r._customizeFeatureA
}

// Set is ShopId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoPutAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r TaobaoUsceslIteminfoPutAPIRequest) GetShopId() int64 {
	return r._shopId
}
