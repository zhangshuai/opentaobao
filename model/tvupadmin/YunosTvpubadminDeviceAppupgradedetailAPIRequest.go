package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 API请求
yunos.tvpubadmin.device.appupgradedetail

获取应用升级详情
*/
type YunosTvpubadminDeviceAppupgradedetailAPIRequest struct {
    model.Params
    // 应用升级的ID
    _versionId   int64
    // 牌照方
    _license   int64
}

// 初始化YunosTvpubadminDeviceAppupgradedetailAPIRequest对象
func NewYunosTvpubadminDeviceAppupgradedetailRequest() *YunosTvpubadminDeviceAppupgradedetailAPIRequest{
    return &YunosTvpubadminDeviceAppupgradedetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.appupgradedetail"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 应用升级的ID
func (r *YunosTvpubadminDeviceAppupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetVersionId() int64 {
    return r._versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceAppupgradedetailAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetLicense() int64 {
    return r._license
}