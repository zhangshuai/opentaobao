package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemUpdateAPIRequest 物流宝商品修改 API请求
// taobao.wlb.item.update
//
// 修改物流宝商品信息
type TaobaoWlbItemUpdateAPIRequest struct {
	model.Params
	// 需要修改的商品属性值的列表，如果属性不存在，则新增属性
	_updatePropertyKeyList string
	// 需要删除的商品属性key列表
	_deletePropertyKeyList string
	// 需要修改的属性值的列表
	_updatePropertyValueList string
	// 要修改的商品id
	_id int64
	// 要修改的商品名称
	_name string
	// 要修改的商品标题
	_title string
	// 要修改的商品备注
	_remark string
	// 是否易碎品
	_isFriable bool
	// 是否危险品
	_isDangerous bool
	// 商品颜色
	_color string
	// 商品重量，单位G
	_weight int64
	// 商品长度，单位厘米
	_length int64
	// 商品宽度，单位厘米
	_width int64
	// 商品高度，单位厘米
	_height int64
	// 商品体积，单位立方厘米
	_volume int64
	// 商品货类
	_goodsCat string
	// 商品计价货类
	_pricingCat string
	// 商品包装材料类型
	_packageMaterial string
}

// NewTaobaoWlbItemUpdateRequest 初始化TaobaoWlbItemUpdateAPIRequest对象
func NewTaobaoWlbItemUpdateRequest() *TaobaoWlbItemUpdateAPIRequest {
	return &TaobaoWlbItemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UpdatePropertyKeyList Setter
// 需要修改的商品属性值的列表，如果属性不存在，则新增属性
func (r *TaobaoWlbItemUpdateAPIRequest) SetUpdatePropertyKeyList(_updatePropertyKeyList string) error {
	r._updatePropertyKeyList = _updatePropertyKeyList
	r.Set("update_property_key_list", _updatePropertyKeyList)
	return nil
}

// Get UpdatePropertyKeyList Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetUpdatePropertyKeyList() string {
	return r._updatePropertyKeyList
}

// Set is DeletePropertyKeyList Setter
// 需要删除的商品属性key列表
func (r *TaobaoWlbItemUpdateAPIRequest) SetDeletePropertyKeyList(_deletePropertyKeyList string) error {
	r._deletePropertyKeyList = _deletePropertyKeyList
	r.Set("delete_property_key_list", _deletePropertyKeyList)
	return nil
}

// Get DeletePropertyKeyList Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetDeletePropertyKeyList() string {
	return r._deletePropertyKeyList
}

// Set is UpdatePropertyValueList Setter
// 需要修改的属性值的列表
func (r *TaobaoWlbItemUpdateAPIRequest) SetUpdatePropertyValueList(_updatePropertyValueList string) error {
	r._updatePropertyValueList = _updatePropertyValueList
	r.Set("update_property_value_list", _updatePropertyValueList)
	return nil
}

// Get UpdatePropertyValueList Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetUpdatePropertyValueList() string {
	return r._updatePropertyValueList
}

// Set is Id Setter
// 要修改的商品id
func (r *TaobaoWlbItemUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetId() int64 {
	return r._id
}

// Set is Name Setter
// 要修改的商品名称
func (r *TaobaoWlbItemUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetName() string {
	return r._name
}

// Set is Title Setter
// 要修改的商品标题
func (r *TaobaoWlbItemUpdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetTitle() string {
	return r._title
}

// Set is Remark Setter
// 要修改的商品备注
func (r *TaobaoWlbItemUpdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetRemark() string {
	return r._remark
}

// Set is IsFriable Setter
// 是否易碎品
func (r *TaobaoWlbItemUpdateAPIRequest) SetIsFriable(_isFriable bool) error {
	r._isFriable = _isFriable
	r.Set("is_friable", _isFriable)
	return nil
}

// Get IsFriable Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetIsFriable() bool {
	return r._isFriable
}

// Set is IsDangerous Setter
// 是否危险品
func (r *TaobaoWlbItemUpdateAPIRequest) SetIsDangerous(_isDangerous bool) error {
	r._isDangerous = _isDangerous
	r.Set("is_dangerous", _isDangerous)
	return nil
}

// Get IsDangerous Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetIsDangerous() bool {
	return r._isDangerous
}

// Set is Color Setter
// 商品颜色
func (r *TaobaoWlbItemUpdateAPIRequest) SetColor(_color string) error {
	r._color = _color
	r.Set("color", _color)
	return nil
}

// Get Color Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetColor() string {
	return r._color
}

// Set is Weight Setter
// 商品重量，单位G
func (r *TaobaoWlbItemUpdateAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// Get Weight Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetWeight() int64 {
	return r._weight
}

// Set is Length Setter
// 商品长度，单位厘米
func (r *TaobaoWlbItemUpdateAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// Get Length Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetLength() int64 {
	return r._length
}

// Set is Width Setter
// 商品宽度，单位厘米
func (r *TaobaoWlbItemUpdateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// Get Width Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetWidth() int64 {
	return r._width
}

// Set is Height Setter
// 商品高度，单位厘米
func (r *TaobaoWlbItemUpdateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is Volume Setter
// 商品体积，单位立方厘米
func (r *TaobaoWlbItemUpdateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// Get Volume Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetVolume() int64 {
	return r._volume
}

// Set is GoodsCat Setter
// 商品货类
func (r *TaobaoWlbItemUpdateAPIRequest) SetGoodsCat(_goodsCat string) error {
	r._goodsCat = _goodsCat
	r.Set("goods_cat", _goodsCat)
	return nil
}

// Get GoodsCat Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetGoodsCat() string {
	return r._goodsCat
}

// Set is PricingCat Setter
// 商品计价货类
func (r *TaobaoWlbItemUpdateAPIRequest) SetPricingCat(_pricingCat string) error {
	r._pricingCat = _pricingCat
	r.Set("pricing_cat", _pricingCat)
	return nil
}

// Get PricingCat Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetPricingCat() string {
	return r._pricingCat
}

// Set is PackageMaterial Setter
// 商品包装材料类型
func (r *TaobaoWlbItemUpdateAPIRequest) SetPackageMaterial(_packageMaterial string) error {
	r._packageMaterial = _packageMaterial
	r.Set("package_material", _packageMaterial)
	return nil
}

// Get PackageMaterial Getter
func (r TaobaoWlbItemUpdateAPIRequest) GetPackageMaterial() string {
	return r._packageMaterial
}
