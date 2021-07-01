package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出所有项目的药物属性和药品信息 API请求
alibaba.alihealth.drugcode.drugfactory.exportattribute

导出所有项目的药物属性和药品信息
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest struct {
    model.Params
    // 企业id
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportattribute"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetRefEntId() string {
    return r._refEntId
}