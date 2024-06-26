package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaMatchAPIRequest product匹配接口 API请求
// tmall.product.schema.match
//
// 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchAPIRequest struct {
	model.Params
	// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
	_propvalues string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// NewTmallProductSchemaMatchRequest 初始化TmallProductSchemaMatchAPIRequest对象
func NewTmallProductSchemaMatchRequest() *TmallProductSchemaMatchAPIRequest {
	return &TmallProductSchemaMatchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallProductSchemaMatchAPIRequest) Reset() {
	r._propvalues = ""
	r._categoryId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSchemaMatchAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSchemaMatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductSchemaMatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPropvalues is Propvalues Setter
// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
func (r *TmallProductSchemaMatchAPIRequest) SetPropvalues(_propvalues string) error {
	r._propvalues = _propvalues
	r.Set("propvalues", _propvalues)
	return nil
}

// GetPropvalues Propvalues Getter
func (r TmallProductSchemaMatchAPIRequest) GetPropvalues() string {
	return r._propvalues
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductSchemaMatchAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallProductSchemaMatchAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

var poolTmallProductSchemaMatchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallProductSchemaMatchRequest()
	},
}

// GetTmallProductSchemaMatchRequest 从 sync.Pool 获取 TmallProductSchemaMatchAPIRequest
func GetTmallProductSchemaMatchAPIRequest() *TmallProductSchemaMatchAPIRequest {
	return poolTmallProductSchemaMatchAPIRequest.Get().(*TmallProductSchemaMatchAPIRequest)
}

// ReleaseTmallProductSchemaMatchAPIRequest 将 TmallProductSchemaMatchAPIRequest 放入 sync.Pool
func ReleaseTmallProductSchemaMatchAPIRequest(v *TmallProductSchemaMatchAPIRequest) {
	v.Reset()
	poolTmallProductSchemaMatchAPIRequest.Put(v)
}
