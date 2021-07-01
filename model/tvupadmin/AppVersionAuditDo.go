package tvupadmin

// AppVersionAuditDo 
type AppVersionAuditDo struct {
    // 主键ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 应用名
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    // 应用包名
    AppPackageName   string `json:"app_package_name,omitempty" xml:"app_package_name,omitempty"`
    // 版本号
    VersionName   string `json:"version_name,omitempty" xml:"version_name,omitempty"`
    // 版本Code
    VersionCode   string `json:"version_code,omitempty" xml:"version_code,omitempty"`
    // 发布说明
    ReleaseNote   string `json:"release_note,omitempty" xml:"release_note,omitempty"`
    // 下载地址
    DownloadUrl   string `json:"download_url,omitempty" xml:"download_url,omitempty"`
    // 安装包大小
    Size   string `json:"size,omitempty" xml:"size,omitempty"`
    // 审核状态
    AuditStatus   string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
    // 审核意见
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModify   string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
    // 审核时间
    GmtAudit   string `json:"gmt_audit,omitempty" xml:"gmt_audit,omitempty"`
}
