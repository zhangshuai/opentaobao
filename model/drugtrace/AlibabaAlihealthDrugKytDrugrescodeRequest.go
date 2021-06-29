package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品码段信息 API请求
alibaba.alihealth.drug.kyt.drugrescode

查询药品码段信息
*/
type AlibabaAlihealthDrugKytDrugrescodeRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 药品通用名
    physicName   string
    // 批准文号
    approvalLicenceNo   string
    // 开始日期
    startDate   string
    // 结束日期
    endDate   string
    // 页大小
    pageSize   int64
    // 页码
    page   int64
    // 企业名称
    entName   string
    // 包装规格
    packageSpec   string
    // 制剂规格
    prepnSpec   string
}

// 初始化AlibabaAlihealthDrugKytDrugrescodeRequest对象
func NewAlibabaAlihealthDrugKytDrugrescodeRequest() *AlibabaAlihealthDrugKytDrugrescodeRequest{
    return &AlibabaAlihealthDrugKytDrugrescodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugrescode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetRefEntId() string {
    return r.refEntId
}
// PhysicName Setter
// 药品通用名
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPhysicName(physicName string) error {
    r.physicName = physicName
    r.Set("physic_name", physicName)
    return nil
}

// PhysicName Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPhysicName() string {
    return r.physicName
}
// ApprovalLicenceNo Setter
// 批准文号
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetApprovalLicenceNo(approvalLicenceNo string) error {
    r.approvalLicenceNo = approvalLicenceNo
    r.Set("approval_licence_no", approvalLicenceNo)
    return nil
}

// ApprovalLicenceNo Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApprovalLicenceNo() string {
    return r.approvalLicenceNo
}
// StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetEndDate() string {
    return r.endDate
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPage() int64 {
    return r.page
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetEntName() string {
    return r.entName
}
// PackageSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPackageSpec(packageSpec string) error {
    r.packageSpec = packageSpec
    r.Set("package_spec", packageSpec)
    return nil
}

// PackageSpec Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPackageSpec() string {
    return r.packageSpec
}
// PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPrepnSpec(prepnSpec string) error {
    r.prepnSpec = prepnSpec
    r.Set("prepn_spec", prepnSpec)
    return nil
}

// PrepnSpec Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPrepnSpec() string {
    return r.prepnSpec
}