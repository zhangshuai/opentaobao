package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSimpleschemaAddAPIRequest 天猫简化发布商品 API请求
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddAPIRequest struct {
	model.Params
	// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
	_schemaXmlFields string
}

// NewTmallItemSimpleschemaAddRequest 初始化TmallItemSimpleschemaAddAPIRequest对象
func NewTmallItemSimpleschemaAddRequest() *TmallItemSimpleschemaAddAPIRequest {
	return &TmallItemSimpleschemaAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSimpleschemaAddAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSimpleschemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
func (r *TmallItemSimpleschemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallItemSimpleschemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

var poolTmallItemSimpleschemaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSimpleschemaAddRequest()
	},
}

// GetTmallItemSimpleschemaAddRequest 从 sync.Pool 获取 TmallItemSimpleschemaAddAPIRequest
func GetTmallItemSimpleschemaAddAPIRequest() *TmallItemSimpleschemaAddAPIRequest {
	return poolTmallItemSimpleschemaAddAPIRequest.Get().(*TmallItemSimpleschemaAddAPIRequest)
}

// ReleaseTmallItemSimpleschemaAddAPIRequest 将 TmallItemSimpleschemaAddAPIRequest 放入 sync.Pool
func ReleaseTmallItemSimpleschemaAddAPIRequest(v *TmallItemSimpleschemaAddAPIRequest) {
	v.Reset()
	poolTmallItemSimpleschemaAddAPIRequest.Put(v)
}
