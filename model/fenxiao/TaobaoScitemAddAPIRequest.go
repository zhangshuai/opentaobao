package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemAddAPIRequest 发布后端商品 API请求
// taobao.scitem.add
//
// 发布后端商品
type TaobaoScitemAddAPIRequest struct {
	model.Params
	// 商品名称
	_itemName string
	// 商家编码
	_outerCode string
	// 0.普通供应链商品 1.供应链组合主商品
	_itemType int64
	// 商品属性格式是  p1:v1,p2:v2,p3:v3
	_properties string
	// 条形码
	_barCode string
	// 仓储商编码
	_wmsCode string
	// 是否易碎 0：不是  1：是
	_isFriable int64
	// 是否危险 0：不是  1：是
	_isDangerous int64
	// 是否是贵重品 0:不是 1：是
	_isCostly int64
	// 是否保质期：0:不是 1：是
	_isWarranty int64
	// 重量 单位：g
	_weight int64
	// 长度 单位：mm
	_length int64
	// 宽 单位：mm
	_width int64
	// 高 单位：mm
	_height int64
	// 体积：立方厘米
	_volume int64
	// 价格 单位：分
	_price int64
	// remark
	_remark string
	// 0:液体，1：粉体，2：固体
	_matterStatus int64
	// 品牌id
	_brandId int64
	// brand_Name
	_brandName string
	// spuId或是cspuid
	_spuId int64
	// 1表示区域销售，0或是空是非区域销售
	_isAreaSale int64
}

// NewTaobaoScitemAddRequest 初始化TaobaoScitemAddAPIRequest对象
func NewTaobaoScitemAddRequest() *TaobaoScitemAddAPIRequest {
	return &TaobaoScitemAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemAddAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemName Setter
// 商品名称
func (r *TaobaoScitemAddAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// Get ItemName Getter
func (r TaobaoScitemAddAPIRequest) GetItemName() string {
	return r._itemName
}

// Set is OuterCode Setter
// 商家编码
func (r *TaobaoScitemAddAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// Get OuterCode Getter
func (r TaobaoScitemAddAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// Set is ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemAddAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// Get ItemType Getter
func (r TaobaoScitemAddAPIRequest) GetItemType() int64 {
	return r._itemType
}

// Set is Properties Setter
// 商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemAddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// Get Properties Getter
func (r TaobaoScitemAddAPIRequest) GetProperties() string {
	return r._properties
}

// Set is BarCode Setter
// 条形码
func (r *TaobaoScitemAddAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// Get BarCode Getter
func (r TaobaoScitemAddAPIRequest) GetBarCode() string {
	return r._barCode
}

// Set is WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemAddAPIRequest) SetWmsCode(_wmsCode string) error {
	r._wmsCode = _wmsCode
	r.Set("wms_code", _wmsCode)
	return nil
}

// Get WmsCode Getter
func (r TaobaoScitemAddAPIRequest) GetWmsCode() string {
	return r._wmsCode
}

// Set is IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemAddAPIRequest) SetIsFriable(_isFriable int64) error {
	r._isFriable = _isFriable
	r.Set("is_friable", _isFriable)
	return nil
}

// Get IsFriable Getter
func (r TaobaoScitemAddAPIRequest) GetIsFriable() int64 {
	return r._isFriable
}

// Set is IsDangerous Setter
// 是否危险 0：不是  1：是
func (r *TaobaoScitemAddAPIRequest) SetIsDangerous(_isDangerous int64) error {
	r._isDangerous = _isDangerous
	r.Set("is_dangerous", _isDangerous)
	return nil
}

// Get IsDangerous Getter
func (r TaobaoScitemAddAPIRequest) GetIsDangerous() int64 {
	return r._isDangerous
}

// Set is IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemAddAPIRequest) SetIsCostly(_isCostly int64) error {
	r._isCostly = _isCostly
	r.Set("is_costly", _isCostly)
	return nil
}

// Get IsCostly Getter
func (r TaobaoScitemAddAPIRequest) GetIsCostly() int64 {
	return r._isCostly
}

// Set is IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemAddAPIRequest) SetIsWarranty(_isWarranty int64) error {
	r._isWarranty = _isWarranty
	r.Set("is_warranty", _isWarranty)
	return nil
}

// Get IsWarranty Getter
func (r TaobaoScitemAddAPIRequest) GetIsWarranty() int64 {
	return r._isWarranty
}

// Set is Weight Setter
// 重量 单位：g
func (r *TaobaoScitemAddAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// Get Weight Getter
func (r TaobaoScitemAddAPIRequest) GetWeight() int64 {
	return r._weight
}

// Set is Length Setter
// 长度 单位：mm
func (r *TaobaoScitemAddAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// Get Length Getter
func (r TaobaoScitemAddAPIRequest) GetLength() int64 {
	return r._length
}

// Set is Width Setter
// 宽 单位：mm
func (r *TaobaoScitemAddAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// Get Width Getter
func (r TaobaoScitemAddAPIRequest) GetWidth() int64 {
	return r._width
}

// Set is Height Setter
// 高 单位：mm
func (r *TaobaoScitemAddAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r TaobaoScitemAddAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemAddAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// Get Volume Getter
func (r TaobaoScitemAddAPIRequest) GetVolume() int64 {
	return r._volume
}

// Set is Price Setter
// 价格 单位：分
func (r *TaobaoScitemAddAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// Get Price Getter
func (r TaobaoScitemAddAPIRequest) GetPrice() int64 {
	return r._price
}

// Set is Remark Setter
// remark
func (r *TaobaoScitemAddAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoScitemAddAPIRequest) GetRemark() string {
	return r._remark
}

// Set is MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemAddAPIRequest) SetMatterStatus(_matterStatus int64) error {
	r._matterStatus = _matterStatus
	r.Set("matter_status", _matterStatus)
	return nil
}

// Get MatterStatus Getter
func (r TaobaoScitemAddAPIRequest) GetMatterStatus() int64 {
	return r._matterStatus
}

// Set is BrandId Setter
// 品牌id
func (r *TaobaoScitemAddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// Get BrandId Getter
func (r TaobaoScitemAddAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// Set is BrandName Setter
// brand_Name
func (r *TaobaoScitemAddAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// Get BrandName Getter
func (r TaobaoScitemAddAPIRequest) GetBrandName() string {
	return r._brandName
}

// Set is SpuId Setter
// spuId或是cspuid
func (r *TaobaoScitemAddAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// Get SpuId Getter
func (r TaobaoScitemAddAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// Set is IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemAddAPIRequest) SetIsAreaSale(_isAreaSale int64) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// Get IsAreaSale Getter
func (r TaobaoScitemAddAPIRequest) GetIsAreaSale() int64 {
	return r._isAreaSale
}
