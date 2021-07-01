package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品资质信息 API请求
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。
*/
type AlibabaAlihealthDrugKytGetdruglicenseAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 药品ID
    _drugId   string
}

// 初始化AlibabaAlihealthDrugKytGetdruglicenseAPIRequest对象
func NewAlibabaAlihealthDrugKytGetdruglicenseRequest() *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest{
    return &AlibabaAlihealthDrugKytGetdruglicenseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getdruglicense"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetDrugId() string {
    return r._drugId
}