package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品发布规则接口 API请求
alibaba.gpu.add.schema.get

获取产品发布规则接口
*/
type AlibabaGpuAddSchemaGetRequest struct {
    model.Params
    // 叶子类目ID
    leafCatId   int64
    // 品牌ID
    brandId   int64
    // 当前用户所在渠道如0代表天猫，8代表淘宝
    providerId   int64
}

// 初始化AlibabaGpuAddSchemaGetRequest对象
func NewAlibabaGpuAddSchemaGetRequest() *AlibabaGpuAddSchemaGetRequest{
    return &AlibabaGpuAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuAddSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.gpu.add.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuAddSchemaGetRequest) SetLeafCatId(leafCatId int64) error {
    r.leafCatId = leafCatId
    r.Set("leaf_cat_id", leafCatId)
    return nil
}

// LeafCatId Getter
func (r AlibabaGpuAddSchemaGetRequest) GetLeafCatId() int64 {
    return r.leafCatId
}
// BrandId Setter
// 品牌ID
func (r *AlibabaGpuAddSchemaGetRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r AlibabaGpuAddSchemaGetRequest) GetBrandId() int64 {
    return r.brandId
}
// ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuAddSchemaGetRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuAddSchemaGetRequest) GetProviderId() int64 {
    return r.providerId
}
