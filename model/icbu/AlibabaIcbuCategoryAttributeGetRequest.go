package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目属性获取 API请求
alibaba.icbu.category.attribute.get

根据类目ID获取系统定义的属性
*/
type AlibabaIcbuCategoryAttributeGetRequest struct {
    model.Params
    // 发布类目id
    catId   int64
}

// 初始化AlibabaIcbuCategoryAttributeGetRequest对象
func NewAlibabaIcbuCategoryAttributeGetRequest() *AlibabaIcbuCategoryAttributeGetRequest{
    return &AlibabaIcbuCategoryAttributeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryAttributeGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.attribute.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryAttributeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id
func (r *AlibabaIcbuCategoryAttributeGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryAttributeGetRequest) GetCatId() int64 {
    return r.catId
}
