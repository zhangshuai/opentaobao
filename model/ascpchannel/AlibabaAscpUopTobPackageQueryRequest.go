package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
B2B包裹查询接口 API请求
alibaba.ascp.uop.tob.package.query

供应链中台TOB包裹查询接口
*/
type AlibabaAscpUopTobPackageQueryRequest struct {
    model.Params
    // 系统自动生成
    packageQueryRequest   *Packagequeryrequest
}

// 初始化AlibabaAscpUopTobPackageQueryRequest对象
func NewAlibabaAscpUopTobPackageQueryRequest() *AlibabaAscpUopTobPackageQueryRequest{
    return &AlibabaAscpUopTobPackageQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTobPackageQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.tob.package.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTobPackageQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageQueryRequest Setter
// 系统自动生成
func (r *AlibabaAscpUopTobPackageQueryRequest) SetPackageQueryRequest(packageQueryRequest *Packagequeryrequest) error {
    r.packageQueryRequest = packageQueryRequest
    r.Set("package_query_request", packageQueryRequest)
    return nil
}

// PackageQueryRequest Getter
func (r AlibabaAscpUopTobPackageQueryRequest) GetPackageQueryRequest() *Packagequeryrequest {
    return r.packageQueryRequest
}