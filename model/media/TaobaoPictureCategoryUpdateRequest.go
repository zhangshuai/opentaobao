package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新图片分类 API请求
taobao.picture.category.update

更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
*/
type TaobaoPictureCategoryUpdateRequest struct {
    model.Params
    // 要更新的图片分类的id
    categoryId   int64
    // 图片分类的新名字，最大长度20字符，不能为空
    categoryName   string
    // 图片分类的新父分类id
    parentId   int64
}

// 初始化TaobaoPictureCategoryUpdateRequest对象
func NewTaobaoPictureCategoryUpdateRequest() *TaobaoPictureCategoryUpdateRequest{
    return &TaobaoPictureCategoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureCategoryUpdateRequest) GetApiMethodName() string {
    return "taobao.picture.category.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureCategoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 要更新的图片分类的id
func (r *TaobaoPictureCategoryUpdateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TaobaoPictureCategoryUpdateRequest) GetCategoryId() int64 {
    return r.categoryId
}
// CategoryName Setter
// 图片分类的新名字，最大长度20字符，不能为空
func (r *TaobaoPictureCategoryUpdateRequest) SetCategoryName(categoryName string) error {
    r.categoryName = categoryName
    r.Set("category_name", categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoPictureCategoryUpdateRequest) GetCategoryName() string {
    return r.categoryName
}
// ParentId Setter
// 图片分类的新父分类id
func (r *TaobaoPictureCategoryUpdateRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

// ParentId Getter
func (r TaobaoPictureCategoryUpdateRequest) GetParentId() int64 {
    return r.parentId
}
