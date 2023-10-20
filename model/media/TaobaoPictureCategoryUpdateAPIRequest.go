package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturecategoryupdateAPIRequest 更新图片分类 API请求
// taobao.picture.category.update
//
// 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaopicturecategoryupdateAPIRequest struct {
	model.Params
	// 图片分类的新名字，最大长度20字符，不能为空
	_categoryName string
	// 要更新的图片分类的id
	_categoryId int64
	// 图片分类的新父分类id
	_parentId int64
}

// NewTaobaopicturecategoryupdateRequest 初始化TaobaopicturecategoryupdateAPIRequest对象
func NewTaobaopicturecategoryupdateRequest() *TaobaopicturecategoryupdateAPIRequest {
	return &TaobaopicturecategoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopicturecategoryupdateAPIRequest) GetApiMethodName() string {
	return "taobao.picture.category.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopicturecategoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopicturecategoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryName is CategoryName Setter
// 图片分类的新名字，最大长度20字符，不能为空
func (r *TaobaopicturecategoryupdateAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// GetCategoryName CategoryName Getter
func (r TaobaopicturecategoryupdateAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// SetCategoryId is CategoryId Setter
// 要更新的图片分类的id
func (r *TaobaopicturecategoryupdateAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaopicturecategoryupdateAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetParentId is ParentId Setter
// 图片分类的新父分类id
func (r *TaobaopicturecategoryupdateAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r TaobaopicturecategoryupdateAPIRequest) GetParentId() int64 {
	return r._parentId
}
