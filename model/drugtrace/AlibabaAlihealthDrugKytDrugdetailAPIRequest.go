package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品详细信息 API请求
alibaba.alihealth.drug.kyt.drugdetail

查询药品详细信息
*/
type AlibabaAlihealthDrugKytDrugdetailAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 药品ID
    _drugEntBaseInfoId   string
}

// 初始化AlibabaAlihealthDrugKytDrugdetailAPIRequest对象
func NewAlibabaAlihealthDrugKytDrugdetailRequest() *AlibabaAlihealthDrugKytDrugdetailAPIRequest{
    return &AlibabaAlihealthDrugKytDrugdetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugdetailAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytDrugdetailAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
    r._drugEntBaseInfoId = _drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetDrugEntBaseInfoId() string {
    return r._drugEntBaseInfoId
}