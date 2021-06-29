package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增网点覆盖的服务 API请求
tmall.servicecenter.servicestore.createservicestorecoverservice

新增网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务已存在，会新增失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
*/
type TmallServicecenterServicestoreCreateservicestorecoverserviceRequest struct {
    model.Params
    // 业务类型
    bizType   string
    // json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
    categoryIdsAndBrandIds   string
    // serviceCodes列表,|分隔
    serviceCodes   string
    // 网点编码
    serviceStoreCode   string
}

// 初始化TmallServicecenterServicestoreCreateservicestorecoverserviceRequest对象
func NewTmallServicecenterServicestoreCreateservicestorecoverserviceRequest() *TmallServicecenterServicestoreCreateservicestorecoverserviceRequest{
    return &TmallServicecenterServicestoreCreateservicestorecoverserviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.createservicestorecoverservice"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetBizType() string {
    return r.bizType
}
// CategoryIdsAndBrandIds Setter
// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) SetCategoryIdsAndBrandIds(categoryIdsAndBrandIds string) error {
    r.categoryIdsAndBrandIds = categoryIdsAndBrandIds
    r.Set("category_ids_and_brand_ids", categoryIdsAndBrandIds)
    return nil
}

// CategoryIdsAndBrandIds Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetCategoryIdsAndBrandIds() string {
    return r.categoryIdsAndBrandIds
}
// ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) SetServiceCodes(serviceCodes string) error {
    r.serviceCodes = serviceCodes
    r.Set("service_codes", serviceCodes)
    return nil
}

// ServiceCodes Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetServiceCodes() string {
    return r.serviceCodes
}
// ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}
