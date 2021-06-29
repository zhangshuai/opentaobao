package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外投广告 API请求
yunos.appstore.open.getads

将广告外投给外部合作伙伴
*/
type YunosAppstoreOpenGetadsRequest struct {
    model.Params
    // 请求id
    rid   string
    // 指定广告分类
    cats   []string
    // 是否排除已安装
    excludeInstall   bool
    // 场景或页面标识
    caseId   string
    // ssp标识
    ssp   string
    // 结算类型
    feeType   string
    // 客户端来源ip
    clientIp   string
    // 广告指定包名
    pkgs   []string
    // 客户端版本号
    clientVerCode   int64
    // 是否映射到uuid
    tryMapToUuid   bool
    // 排除包名列表
    excludePkgs   []string
    // 设备唯一标识
    deviceId   string
    // 广告数量
    size   int64
    // 排除分类
    excludeCats   []string
    // 创意模板id列表
    templateIds   []int64
    // 广告底价
    mrp   int64
    // 请求特征集
    options   int64
}

// 初始化YunosAppstoreOpenGetadsRequest对象
func NewYunosAppstoreOpenGetadsRequest() *YunosAppstoreOpenGetadsRequest{
    return &YunosAppstoreOpenGetadsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAppstoreOpenGetadsRequest) GetApiMethodName() string {
    return "yunos.appstore.open.getads"
}

// IRequest interface 方法, 获取API参数
func (r YunosAppstoreOpenGetadsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rid Setter
// 请求id
func (r *YunosAppstoreOpenGetadsRequest) SetRid(rid string) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

// Rid Getter
func (r YunosAppstoreOpenGetadsRequest) GetRid() string {
    return r.rid
}
// Cats Setter
// 指定广告分类
func (r *YunosAppstoreOpenGetadsRequest) SetCats(cats []string) error {
    r.cats = cats
    r.Set("cats", cats)
    return nil
}

// Cats Getter
func (r YunosAppstoreOpenGetadsRequest) GetCats() []string {
    return r.cats
}
// ExcludeInstall Setter
// 是否排除已安装
func (r *YunosAppstoreOpenGetadsRequest) SetExcludeInstall(excludeInstall bool) error {
    r.excludeInstall = excludeInstall
    r.Set("exclude_install", excludeInstall)
    return nil
}

// ExcludeInstall Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludeInstall() bool {
    return r.excludeInstall
}
// CaseId Setter
// 场景或页面标识
func (r *YunosAppstoreOpenGetadsRequest) SetCaseId(caseId string) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

// CaseId Getter
func (r YunosAppstoreOpenGetadsRequest) GetCaseId() string {
    return r.caseId
}
// Ssp Setter
// ssp标识
func (r *YunosAppstoreOpenGetadsRequest) SetSsp(ssp string) error {
    r.ssp = ssp
    r.Set("ssp", ssp)
    return nil
}

// Ssp Getter
func (r YunosAppstoreOpenGetadsRequest) GetSsp() string {
    return r.ssp
}
// FeeType Setter
// 结算类型
func (r *YunosAppstoreOpenGetadsRequest) SetFeeType(feeType string) error {
    r.feeType = feeType
    r.Set("fee_type", feeType)
    return nil
}

// FeeType Getter
func (r YunosAppstoreOpenGetadsRequest) GetFeeType() string {
    return r.feeType
}
// ClientIp Setter
// 客户端来源ip
func (r *YunosAppstoreOpenGetadsRequest) SetClientIp(clientIp string) error {
    r.clientIp = clientIp
    r.Set("client_ip", clientIp)
    return nil
}

// ClientIp Getter
func (r YunosAppstoreOpenGetadsRequest) GetClientIp() string {
    return r.clientIp
}
// Pkgs Setter
// 广告指定包名
func (r *YunosAppstoreOpenGetadsRequest) SetPkgs(pkgs []string) error {
    r.pkgs = pkgs
    r.Set("pkgs", pkgs)
    return nil
}

// Pkgs Getter
func (r YunosAppstoreOpenGetadsRequest) GetPkgs() []string {
    return r.pkgs
}
// ClientVerCode Setter
// 客户端版本号
func (r *YunosAppstoreOpenGetadsRequest) SetClientVerCode(clientVerCode int64) error {
    r.clientVerCode = clientVerCode
    r.Set("client_ver_code", clientVerCode)
    return nil
}

// ClientVerCode Getter
func (r YunosAppstoreOpenGetadsRequest) GetClientVerCode() int64 {
    return r.clientVerCode
}
// TryMapToUuid Setter
// 是否映射到uuid
func (r *YunosAppstoreOpenGetadsRequest) SetTryMapToUuid(tryMapToUuid bool) error {
    r.tryMapToUuid = tryMapToUuid
    r.Set("try_map_to_uuid", tryMapToUuid)
    return nil
}

// TryMapToUuid Getter
func (r YunosAppstoreOpenGetadsRequest) GetTryMapToUuid() bool {
    return r.tryMapToUuid
}
// ExcludePkgs Setter
// 排除包名列表
func (r *YunosAppstoreOpenGetadsRequest) SetExcludePkgs(excludePkgs []string) error {
    r.excludePkgs = excludePkgs
    r.Set("exclude_pkgs", excludePkgs)
    return nil
}

// ExcludePkgs Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludePkgs() []string {
    return r.excludePkgs
}
// DeviceId Setter
// 设备唯一标识
func (r *YunosAppstoreOpenGetadsRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r YunosAppstoreOpenGetadsRequest) GetDeviceId() string {
    return r.deviceId
}
// Size Setter
// 广告数量
func (r *YunosAppstoreOpenGetadsRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

// Size Getter
func (r YunosAppstoreOpenGetadsRequest) GetSize() int64 {
    return r.size
}
// ExcludeCats Setter
// 排除分类
func (r *YunosAppstoreOpenGetadsRequest) SetExcludeCats(excludeCats []string) error {
    r.excludeCats = excludeCats
    r.Set("exclude_cats", excludeCats)
    return nil
}

// ExcludeCats Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludeCats() []string {
    return r.excludeCats
}
// TemplateIds Setter
// 创意模板id列表
func (r *YunosAppstoreOpenGetadsRequest) SetTemplateIds(templateIds []int64) error {
    r.templateIds = templateIds
    r.Set("template_ids", templateIds)
    return nil
}

// TemplateIds Getter
func (r YunosAppstoreOpenGetadsRequest) GetTemplateIds() []int64 {
    return r.templateIds
}
// Mrp Setter
// 广告底价
func (r *YunosAppstoreOpenGetadsRequest) SetMrp(mrp int64) error {
    r.mrp = mrp
    r.Set("mrp", mrp)
    return nil
}

// Mrp Getter
func (r YunosAppstoreOpenGetadsRequest) GetMrp() int64 {
    return r.mrp
}
// Options Setter
// 请求特征集
func (r *YunosAppstoreOpenGetadsRequest) SetOptions(options int64) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r YunosAppstoreOpenGetadsRequest) GetOptions() int64 {
    return r.options
}
