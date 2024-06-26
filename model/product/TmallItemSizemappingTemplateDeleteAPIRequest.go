package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplateDeleteAPIRequest 删除天猫商品尺码表模板 API请求
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// NewTmallItemSizemappingTemplateDeleteRequest 初始化TmallItemSizemappingTemplateDeleteAPIRequest对象
func NewTmallItemSizemappingTemplateDeleteRequest() *TmallItemSizemappingTemplateDeleteAPIRequest {
	return &TmallItemSizemappingTemplateDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSizemappingTemplateDeleteAPIRequest) Reset() {
	r._templateId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateDeleteAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetTemplateId() int64 {
	return r._templateId
}

var poolTmallItemSizemappingTemplateDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSizemappingTemplateDeleteRequest()
	},
}

// GetTmallItemSizemappingTemplateDeleteRequest 从 sync.Pool 获取 TmallItemSizemappingTemplateDeleteAPIRequest
func GetTmallItemSizemappingTemplateDeleteAPIRequest() *TmallItemSizemappingTemplateDeleteAPIRequest {
	return poolTmallItemSizemappingTemplateDeleteAPIRequest.Get().(*TmallItemSizemappingTemplateDeleteAPIRequest)
}

// ReleaseTmallItemSizemappingTemplateDeleteAPIRequest 将 TmallItemSizemappingTemplateDeleteAPIRequest 放入 sync.Pool
func ReleaseTmallItemSizemappingTemplateDeleteAPIRequest(v *TmallItemSizemappingTemplateDeleteAPIRequest) {
	v.Reset()
	poolTmallItemSizemappingTemplateDeleteAPIRequest.Put(v)
}
