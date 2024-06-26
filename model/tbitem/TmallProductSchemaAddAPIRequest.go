package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaAddAPIRequest 使用Schema文件发布一个产品 API请求
// tmall.product.schema.add
//
// Schema体系发布一个产品
type TmallProductSchemaAddAPIRequest struct {
	model.Params
	// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
	_xmlData string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
}

// NewTmallProductSchemaAddRequest 初始化TmallProductSchemaAddAPIRequest对象
func NewTmallProductSchemaAddRequest() *TmallProductSchemaAddAPIRequest {
	return &TmallProductSchemaAddAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallProductSchemaAddAPIRequest) Reset() {
	r._xmlData = ""
	r._categoryId = 0
	r._brandId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSchemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSchemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductSchemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
func (r *TmallProductSchemaAddAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallProductSchemaAddAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductSchemaAddAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallProductSchemaAddAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *TmallProductSchemaAddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallProductSchemaAddAPIRequest) GetBrandId() int64 {
	return r._brandId
}

var poolTmallProductSchemaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallProductSchemaAddRequest()
	},
}

// GetTmallProductSchemaAddRequest 从 sync.Pool 获取 TmallProductSchemaAddAPIRequest
func GetTmallProductSchemaAddAPIRequest() *TmallProductSchemaAddAPIRequest {
	return poolTmallProductSchemaAddAPIRequest.Get().(*TmallProductSchemaAddAPIRequest)
}

// ReleaseTmallProductSchemaAddAPIRequest 将 TmallProductSchemaAddAPIRequest 放入 sync.Pool
func ReleaseTmallProductSchemaAddAPIRequest(v *TmallProductSchemaAddAPIRequest) {
	v.Reset()
	poolTmallProductSchemaAddAPIRequest.Put(v)
}
