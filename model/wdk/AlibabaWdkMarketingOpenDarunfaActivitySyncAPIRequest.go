package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动数据同步 API请求
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest struct {
    model.Params
    // 大润发活动数据
    _activityList   []DrfTxdActivityBo
    // 门店Id
    _shopId   string
    // 版本ID
    _versionId   int64
}

// 初始化AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest对象
func NewAlibabaWdkMarketingOpenDarunfaActivitySyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest{
    return &AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.darunfa.activity.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityList Setter
// 大润发活动数据
func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) SetActivityList(_activityList []DrfTxdActivityBo) error {
    r._activityList = _activityList
    r.Set("activity_list", _activityList)
    return nil
}

// ActivityList Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) GetActivityList() []DrfTxdActivityBo {
    return r._activityList
}
// ShopId Setter
// 门店Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) GetShopId() string {
    return r._shopId
}
// VersionId Setter
// 版本ID
func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest) GetVersionId() int64 {
    return r._versionId
}